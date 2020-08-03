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

package manager

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/northwesternmutual/grammes/gremconnect"
	"github.com/northwesternmutual/grammes/logging"
)

// MOCKDIALER

type mockDialer gremconnect.WebSocket

func (*mockDialer) Connect() error                    { return connect }
func (*mockDialer) Close() error                      { return nil }
func (*mockDialer) Write([]byte) error                { return nil }
func (m *mockDialer) Read() ([]byte, error)           { return []byte(response), nil }
func (*mockDialer) Ping(chan error)                   {}
func (*mockDialer) IsConnected() bool                 { return isConnected }
func (*mockDialer) IsDisposed() bool                  { return isDisposed }
func (*mockDialer) Auth() (*gremconnect.Auth, error)  { return &gremconnect.Auth{}, nil }
func (*mockDialer) Address() string                   { return "" }
func (m *mockDialer) GetQuit() chan struct{}          { return make(chan struct{}) }
func (*mockDialer) SetAuth(string, string)            {}
func (*mockDialer) SetTimeout(time.Duration)          {}
func (*mockDialer) SetPingInterval(time.Duration)     {}
func (*mockDialer) SetWritingWait(time.Duration)      {}
func (*mockDialer) SetReadingWait(time.Duration)      {}
func (*mockDialer) SetWriteBufferSize(int)            {}
func (*mockDialer) SetReadBufferSize(int)             {}
func (*mockDialer) SetHandshakeTimeout(time.Duration) {}
func (*mockDialer) SetCompression(bool)               {}

func TestSetLoggerQM(t *testing.T) {
	Convey("Given a dialer, string executor and query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([][]byte, error) { return nil, nil }
		qm := newQueryManager(dialer, logging.NewNilLogger(), execute)
		Convey("When setLogger is called we should not encounter any errors", func() {
			qm.setLogger(logging.NewNilLogger())
		})
	})
}

func TestExecuteQuery(t *testing.T) {
	Convey("Given a dialer, string executor and query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([][]byte, error) { return nil, nil }
		qm := newQueryManager(dialer, logging.NewNilLogger(), execute)
		Convey("When ExecuteQuery is called", func() {
			var q mockQuery
			_, err := qm.ExecuteQuery(q)
			Convey("Then the returned error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestExecuteStringQuery(t *testing.T) {
	Convey("Given a dialer, string executor and query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([][]byte, error) { return nil, nil }
		qm := newQueryManager(dialer, logging.NewNilLogger(), execute)
		Convey("When ExecuteStringQuery is called", func() {
			_, err := qm.ExecuteStringQuery("testquery")
			Convey("Then the returned error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestExecuteBoundQuery(t *testing.T) {
	Convey("Given a dialer, string executor and query manager", t, func() {
		dialer := gremconnect.NewWebSocketDialer("testaddress")
		execute := func(string, map[string]string, map[string]string) ([][]byte, error) { return nil, nil }
		qm := newQueryManager(dialer, logging.NewNilLogger(), execute)
		Convey("When ExecuteBoundQuery is called", func() {
			var q mockQuery
			var b, r map[string]string
			_, err := qm.ExecuteBoundQuery(q, b, r)
			Convey("Then the returned error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestExecuteBoundStringQueryDisposedConnection(t *testing.T) {
	Convey("Given a dialer, string executor and query manager", t, func() {
		dialer := &mockDialer{}
		execute := func(string, map[string]string, map[string]string) ([][]byte, error) { return nil, nil }
		qm := newQueryManager(dialer, logging.NewNilLogger(), execute)
		Convey("When ExecuteBoundStringQuery is called with a disposed connection", func() {
			var b, r map[string]string
			isDisposed = true
			_, err := qm.ExecuteBoundStringQuery("testquery", b, r)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
