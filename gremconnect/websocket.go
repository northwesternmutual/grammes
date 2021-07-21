// Copyright (c) 2018 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gremconnect

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WebSocket will hold all of the data used
// to dial to the gremlin server and sustain
// a stable connection by pinging it regularly.
type WebSocket struct {
	address           string
	conn              *websocket.Conn
	auth              *Auth
	httpAuth          *HTTPAuth
	disposed          bool
	connected         bool
	enableCompression bool
	pingInterval      time.Duration
	writingWait       time.Duration
	readingWait       time.Duration
	timeout           time.Duration
	handshakeTimeout  time.Duration
	writeBufferSize   int
	readBufferSize    int
	Quit              chan struct{}

	sync.RWMutex
}

// Connect will setup the gorilla websocket and
// other configurations to establish a connection
// to the given address.
func (ws *WebSocket) Connect() error {
	var err error
	dialer := websocket.Dialer{
		WriteBufferSize:   ws.writeBufferSize,
		ReadBufferSize:    ws.readBufferSize,
		HandshakeTimeout:  ws.handshakeTimeout,
		EnableCompression: ws.enableCompression,
	}

	// Check if the host address already has the proper
	// /gremlin endpoint at the end of it. If it doesn't
	// then concatenate it to the end of the string.
	// https://groups.google.com/forum/#!msg/gremlin-users/x4hiHsmTsHM/Xe4GcPtRCAAJ
	if !strings.HasSuffix(ws.address, "/gremlin") {
		ws.address = ws.address + "/gremlin"
	}

	// This is a minor hack, but a nice hook-point. We are using the dialer's proxy callback in order to modify the request before it's
	// being sent. That way we can add a custom authentication provider (Like using Amazon Neptune's v4 signer)
	if ws.httpAuth != nil && ws.httpAuth.authProvider != nil {
		dialer.Proxy = func(request *http.Request) (*url.URL, error) {
			return nil, ws.httpAuth.authProvider(request)
		}
	}

	var httpResponse *http.Response
	ws.conn, httpResponse, err = dialer.Dial(ws.address, http.Header{})
	if err != nil {
		if httpResponse != nil {
			//noinspection GoUnhandledErrorResult
			defer httpResponse.Body.Close()

			// Try to read the http response to add context to the error
			errorOutput, readErr := ioutil.ReadAll(httpResponse.Body)
			if readErr == nil {
				return fmt.Errorf("error connecting to address. response: %s. error %v", string(errorOutput), err)
			}
		}

		return err
	}

	ws.connected = true

	handler := func(appData string) error {
		ws.Lock()
		ws.connected = true
		ws.Unlock()
		return nil
	}

	ws.conn.SetPongHandler(handler)

	return nil
}

// IsConnected returns whether the given
// websocket has an established connection.
func (ws *WebSocket) IsConnected() bool {
	return ws.connected
}

// IsDisposed returns whether the given
// websocket has been disposed of its use.
func (ws *WebSocket) IsDisposed() bool {
	return ws.disposed
}

// Write uses the gorilla function to write
// a Binary message to the established connection.
func (ws *WebSocket) Write(msg []byte) error {
	return ws.conn.WriteMessage(websocket.BinaryMessage, msg)
}

// Read uses the gorilla function to read a response
// from the established connection.
func (ws *WebSocket) Read() (msg []byte, err error) {
	_, msg, err = ws.conn.ReadMessage()
	return
}

// Close disposes the websocket and closes the quit
// channel to signal the websocket's ping selection.
func (ws *WebSocket) Close() error {
	defer func() {
		close(ws.Quit) // close the channel to notify our pinger.
		ws.disposed = true
		ws.conn.Close()
	}()

	// Send the server the message that we've closed
	// the connection.
	return ws.conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
}

// Auth returns the websocket's authentication
// information if it's on a secure connection.
func (ws *WebSocket) Auth() (*Auth, error) {
	if ws.auth == nil {
		return nil, errors.New("must create a secure dialer for authentication with the server")
	}

	return ws.auth, nil
}

// Address returns the host address used to
// establish a connection to.
func (ws *WebSocket) Address() string {
	return ws.address
}

// GetQuit returns the quit channel so the websocket
// can communicate to the client that the connection
// has quit.
func (ws *WebSocket) GetQuit() chan struct{} {
	return ws.Quit
}

// Ping runs a routine ping check to the established
// connection and sends error channel a signal if there's
// a detected error if not/how the server responds.
func (ws *WebSocket) Ping(errs chan error) {
	ticker := time.NewTicker(ws.pingInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			connected := true
			// Send a pinging message with the timeout given
			// to the websocket. If there's an error then we lost
			// connection.
			if err := ws.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(ws.writingWait)); err != nil {
				errs <- err
				connected = false
			}
			ws.Lock()
			ws.connected = connected
			ws.Unlock()
		case <-ws.Quit:
			return // Stop pinging if quit.
		}
	}
}

// Configuration functions

// SetAuth will set the authentication to this user and pass
func (ws *WebSocket) SetAuth(user, pass string) {
	ws.auth = &Auth{Username: user, Password: pass}
}

// SetHTTPAuth will set the HTTP authentication provider to this one
func (ws *WebSocket) SetHTTPAuth(provider AuthProvider) {
	ws.httpAuth = &HTTPAuth{provider}
}

// SetTimeout will set the dialing timeout
func (ws *WebSocket) SetTimeout(interval time.Duration) {
	ws.timeout = interval
}

// SetPingInterval sets how often the websocket will ping the server.
func (ws *WebSocket) SetPingInterval(interval time.Duration) {
	ws.pingInterval = interval
}

// SetWritingWait sets how long the wait is for waiting
func (ws *WebSocket) SetWritingWait(interval time.Duration) {
	ws.writingWait = interval
}

// SetReadingWait sets how long the reading will wait
func (ws *WebSocket) SetReadingWait(interval time.Duration) {
	ws.readingWait = interval
}

func (ws *WebSocket) SetWriteBufferSize(writeBufferSize int) {
	ws.writeBufferSize = writeBufferSize
}

func (ws *WebSocket) SetReadBufferSize(readBufferSize int) {
	ws.readBufferSize = readBufferSize
}

func (ws *WebSocket) SetHandshakeTimeout(handshakeTimeout time.Duration) {
	ws.handshakeTimeout = handshakeTimeout
}

func (ws *WebSocket) SetCompression(enableCompression bool) {
	ws.enableCompression = enableCompression
}
