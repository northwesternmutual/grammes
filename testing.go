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

package grammes

import (
	"crypto/tls"
	"errors"
	"time"

	"github.com/gorilla/websocket"

	"github.com/northwesternmutual/grammes/gremconnect"
	"github.com/northwesternmutual/grammes/query/graph"
)

var (
	newVertexResponse = `
	{
		"requestId": "61616161-6161-6161-2d61-6161612d6161",
		"status": {
			"message": "",
			"code": 200,
			"attributes": {}
		},
		"result": {
			"data": [{
				"@type": "g:Vertex",
				"@value": {
					"id": {
						"@type": "g:Int64",
						"@value": 28720
					},
					"label": "newvertex"
				}
			}],
			"meta": {}
		}
	}
	`
	newVertexResponse407 = `
	{
		"requestId": "61616161-6161-6161-2d61-6161612d6161",
		"status": {
			"message": "",
			"code": 407,
			"attributes": {}
		},
		"result": {
			"data": [{
				"@type": "g:Vertex",
				"@value": {
					"id": {
						"@type": "g:Int64",
						"@value": 28720
					},
					"label": "newvertex"
				}
			}],
			"meta": {}
		}
	}
	`
	badResponse = `
	{
		"BADRESPONSE"
	}
	`
)

// TESTLOGGER

type testLogger struct{}

func (testLogger) PrintQuery(string)                    {}
func (testLogger) Debug(string, map[string]interface{}) {}
func (testLogger) Error(string, error)                  {}
func (testLogger) Fatal(string, error)                  {}

type mockString graph.String

func (mockString) String() string { return "TEST" }

type mockDialerStruct struct {
	connect      error
	isConnected  bool
	isDisposed   bool
	response     string
	logger       testLogger
	address      string
	conn         *websocket.Conn
	tlsConfig    *tls.Config
	auth         *gremconnect.Auth
	disposed     bool
	connected    bool
	pingInterval time.Duration
	writingWait  time.Duration
	readingWait  time.Duration
	timeout      time.Duration
	quit         chan struct{}
}

func (m *mockDialerStruct) Connect() error     { return m.connect }
func (*mockDialerStruct) Close() error         { return nil }
func (m *mockDialerStruct) Write([]byte) error { return nil }
func (m *mockDialerStruct) Read() ([]byte, error) {
	time.Sleep(100 * time.Millisecond)
	m.quit <- struct{}{}
	return []byte(m.response), nil
}
func (*mockDialerStruct) Ping(chan error)                  {}
func (m *mockDialerStruct) IsConnected() bool              { return m.isConnected }
func (m *mockDialerStruct) IsDisposed() bool               { return m.isDisposed }
func (*mockDialerStruct) Auth() (*gremconnect.Auth, error) { return &gremconnect.Auth{}, nil }
func (*mockDialerStruct) Address() string                  { return "" }
func (m *mockDialerStruct) GetQuit() chan struct{} {
	m.quit = make(chan struct{})
	return m.quit
}
func (*mockDialerStruct) SetAuth(string, string)        {}
func (*mockDialerStruct) SetTimeout(time.Duration)      {}
func (*mockDialerStruct) SetPingInterval(time.Duration) {}
func (*mockDialerStruct) SetWritingWait(time.Duration)  {}
func (*mockDialerStruct) SetReadingWait(time.Duration)  {}
func (*mockDialerStruct) SetTLSConfig(*tls.Config)      {}

func mockDial(conn gremconnect.Dialer, cfgs ...ClientConfiguration) (*Client, error) {
	c := setupClient()
	c.conn = conn
	for _, conf := range cfgs {
		conf(c)
	}
	return c, nil
}

type mockDialerWriteError gremconnect.WebSocket

func (*mockDialerWriteError) Connect() error                   { return nil }
func (*mockDialerWriteError) Close() error                     { return nil }
func (*mockDialerWriteError) Write([]byte) error               { return errors.New("ERROR") }
func (*mockDialerWriteError) Read() ([]byte, error)            { return nil, nil }
func (*mockDialerWriteError) Ping(chan error)                  {}
func (*mockDialerWriteError) IsConnected() bool                { return true }
func (*mockDialerWriteError) IsDisposed() bool                 { return false }
func (*mockDialerWriteError) Auth() (*gremconnect.Auth, error) { return &gremconnect.Auth{}, nil }
func (*mockDialerWriteError) Address() string                  { return "" }
func (m *mockDialerWriteError) GetQuit() chan struct{} {
	m.Quit = make(chan struct{})
	return m.Quit
}
func (*mockDialerWriteError) SetAuth(string, string)        {}
func (*mockDialerWriteError) SetTimeout(time.Duration)      {}
func (*mockDialerWriteError) SetPingInterval(time.Duration) {}
func (*mockDialerWriteError) SetWritingWait(time.Duration)  {}
func (*mockDialerWriteError) SetReadingWait(time.Duration)  {}
func (*mockDialerWriteError) SetTLSConfig(*tls.Config)      {}

type mockDialerAuthError gremconnect.WebSocket

func (*mockDialerAuthError) Connect() error     { return nil }
func (*mockDialerAuthError) Close() error       { return nil }
func (*mockDialerAuthError) Write([]byte) error { return nil }
func (m *mockDialerAuthError) Read() ([]byte, error) {
	return []byte(newVertexResponse), nil
}
func (*mockDialerAuthError) Ping(chan error)   {}
func (*mockDialerAuthError) IsConnected() bool { return true }
func (*mockDialerAuthError) IsDisposed() bool  { return false }
func (*mockDialerAuthError) Auth() (*gremconnect.Auth, error) {
	return &gremconnect.Auth{}, errors.New("ERROR")
}
func (*mockDialerAuthError) Address() string { return "" }
func (m *mockDialerAuthError) GetQuit() chan struct{} {
	m.Quit = make(chan struct{})
	return m.Quit
}
func (*mockDialerAuthError) SetAuth(string, string)        {}
func (*mockDialerAuthError) SetTimeout(time.Duration)      {}
func (*mockDialerAuthError) SetPingInterval(time.Duration) {}
func (*mockDialerAuthError) SetWritingWait(time.Duration)  {}
func (*mockDialerAuthError) SetReadingWait(time.Duration)  {}
func (*mockDialerAuthError) SetTLSConfig(*tls.Config)      {}

type mockDialerReadError gremconnect.WebSocket

func (*mockDialerReadError) Connect() error     { return nil }
func (*mockDialerReadError) Close() error       { return nil }
func (*mockDialerReadError) Write([]byte) error { return nil }
func (m *mockDialerReadError) Read() ([]byte, error) {
	time.Sleep(100 * time.Millisecond)
	m.Quit <- struct{}{}
	return []byte(newVertexResponse), errors.New("ERROR")
}
func (*mockDialerReadError) Ping(chan error)                  {}
func (*mockDialerReadError) IsConnected() bool                { return true }
func (*mockDialerReadError) IsDisposed() bool                 { return false }
func (*mockDialerReadError) Auth() (*gremconnect.Auth, error) { return &gremconnect.Auth{}, nil }
func (*mockDialerReadError) Address() string                  { return "" }
func (m *mockDialerReadError) GetQuit() chan struct{} {
	m.Quit = make(chan struct{})
	return m.Quit
}
func (*mockDialerReadError) SetAuth(string, string)        {}
func (*mockDialerReadError) SetTimeout(time.Duration)      {}
func (*mockDialerReadError) SetPingInterval(time.Duration) {}
func (*mockDialerReadError) SetWritingWait(time.Duration)  {}
func (*mockDialerReadError) SetReadingWait(time.Duration)  {}
func (*mockDialerReadError) SetTLSConfig(*tls.Config)      {}
