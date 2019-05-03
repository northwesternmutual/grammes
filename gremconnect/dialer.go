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

import "time"

// Dialer will be used to dial in a connection
// between the client and gremlin server without
// directly using a websocket. This leaves room
// for mocking and altered methods of connecting
// to the server.
type Dialer interface {
	// Actions
	Connect() error
	Close() error
	Write(msg []byte) error
	Read() (msg []byte, err error)
	Ping(errs chan error)

	// Checkers
	IsConnected() bool
	IsDisposed() bool

	// Getters
	Auth() (*Auth, error)
	Address() string
	GetQuit() chan struct{}

	// Configuration Setters
	SetAuth(username string, password string)
	SetTimeout(interval time.Duration)
	SetPingInterval(interval time.Duration)
	SetWritingWait(interval time.Duration)
	SetReadingWait(interval time.Duration)
}

// NewWebSocketDialer returns a new WebSocket dialer to use when
// establishing a connection to the Gremlin server. This
// function also assigns default values to the websocket
// if they're not assigned by DialerConfig functions.
func NewWebSocketDialer(address string) Dialer {
	return &WebSocket{
		timeout:      5 * time.Second,
		pingInterval: 60 * time.Second,
		writingWait:  15 * time.Second,
		readingWait:  15 * time.Second,
		connected:    false,
		address:      address,
		Quit:         make(chan struct{}),
	}
}
