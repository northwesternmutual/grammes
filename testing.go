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
	"time"

	"github.com/northwesternmutual/grammes/gremconnect"
	"github.com/northwesternmutual/grammes/query/graph"
)

var (
	readCount         int
	connect           error
	isConnected       = true
	isDisposed        = false
	response          = newVertexResponse
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
	vertexIDsResponse = `
	{
		"requestId": "61616161-6161-6161-2d61-6161612d6161",
		"status": {
			"message": "",
			"code": 200,
			"attributes": {}
		},
		"result": {
			"data": [{
				"@type": "g:Int64",
				"@value": 28720
			}],
			"meta": {}
		}
	}
	`
	vertexCountResponse = `
	{
		"requestId": "61616161-6161-6161-2d61-6161612d6161",
		"status": {
			"message": "",
			"code": 200,
			"attributes": {}
		},
		"result": {
			"data": [{
				"@type": "g:Int64",
				"@value": 1
			}],
			"meta": {}
		}
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

// MOCKDIALER

type mockDialer gremconnect.WebSocket

func (*mockDialer) Connect() error     { return connect }
func (*mockDialer) Close() error       { return nil }
func (*mockDialer) Write([]byte) error { return nil }
func (m *mockDialer) Read() ([]byte, error) {
	if readCount < 1 {
		time.Sleep(100 * time.Millisecond)
		readCount++
		m.Quit <- struct{}{}
		return []byte(response), nil
	}
	return nil, nil
}
func (*mockDialer) Ping(chan error)                  {}
func (*mockDialer) IsConnected() bool                { return isConnected }
func (*mockDialer) IsDisposed() bool                 { return isDisposed }
func (*mockDialer) Auth() (*gremconnect.Auth, error) { return &gremconnect.Auth{}, nil }
func (*mockDialer) Address() string                  { return "" }
func (m *mockDialer) GetQuit() chan struct{} {
	m.Quit = make(chan struct{})
	return m.Quit
}
func (*mockDialer) SetAuth(string, string) {}
func (*mockDialer) SetTimeout(int)         {}
func (*mockDialer) SetPingInterval(int)    {}
func (*mockDialer) SetWritingWait(int)     {}
func (*mockDialer) SetReadingWait(int)     {}
