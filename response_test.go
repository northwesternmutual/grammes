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
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/northwesternmutual/grammes/gremconnect"
)

type mockDialerReadError gremconnect.WebSocket

func (*mockDialerReadError) Connect() error     { return connect }
func (*mockDialerReadError) Close() error       { return nil }
func (*mockDialerReadError) Write([]byte) error { return nil }
func (m *mockDialerReadError) Read() ([]byte, error) {
	if readCount < 1 {
		readCount++
		m.Quit <- struct{}{}
		return []byte(response), errors.New("ERROR")
	}
	return nil, nil
}
func (*mockDialerReadError) Ping(chan error)                  {}
func (*mockDialerReadError) IsConnected() bool                { return isConnected }
func (*mockDialerReadError) IsDisposed() bool                 { return isDisposed }
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

func TestReadWorkerErrorReading(t *testing.T) {
	readCount = 1
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	Convey("Given a client that represents the Gremlin client", t, func() {
		connect = nil
		dialer := &mockDialerReadError{}
		c, _ := Dial(dialer)
		Convey("When there is an error reading the message", func() {
			errReceived := false
			go func() {
				for {
					select {
					case <-c.err:
						errReceived = true
						return
					default:
						continue
					}
				}
			}()
			readCount = 0
			time.Sleep(10 * time.Millisecond)
			Convey("Then the error should be sent through the channel", func() {
				So(errReceived, ShouldBeTrue)
			})
		})
	})
}

func TestReadWorkerErrorHandlingResponse(t *testing.T) {
	readCount = 1
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
		gremMarshalResponse = gremconnect.MarshalResponse
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	gremMarshalResponse = func([]byte) (gremconnect.Response, error) { return gremconnect.Response{}, errors.New("ERROR") }
	Convey("Given a client that represents the Gremlin client", t, func() {
		connect = nil
		dialer := &mockDialer{}
		c, _ := Dial(dialer)
		Convey("When there is an error handling the response", func() {
			errReceived := false
			go func() {
				for {
					select {
					case <-c.err:
						errReceived = true
						return
					default:
						continue
					}
				}
			}()
			readCount = 0
			time.Sleep(125 * time.Millisecond)
			Convey("Then the error should be sent through the channel", func() {
				So(errReceived, ShouldBeTrue)
			})
		})
	})
}

func TestHandleResponseErrorMarshalling(t *testing.T) {
	readCount = 1
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
		gremMarshalResponse = gremconnect.MarshalResponse
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	gremMarshalResponse = func([]byte) (gremconnect.Response, error) { return gremconnect.Response{}, errors.New("ERROR") }
	Convey("Given a client that represents the Gremlin client", t, func() {
		connect = nil
		dialer := &mockDialer{}
		c, _ := Dial(dialer)
		Convey("When handleResponse is called and marshalling throws an error", func() {
			err := c.handleResponse([]byte("testauth"))
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestHandleResponse407Status(t *testing.T) {
	readCount = 1
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
		gremMarshalResponse = gremconnect.MarshalResponse
	}()
	marshalResponse := gremconnect.Response{
		Code: 407,
	}
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	gremMarshalResponse = func([]byte) (gremconnect.Response, error) { return marshalResponse, nil }
	Convey("Given a client that represents the Gremlin client", t, func() {
		connect = nil
		dialer := &mockDialer{}
		c, _ := Dial(dialer)
		Convey("When handleResponse is called and we receive a 407 status code", func() {
			err := c.handleResponse([]byte("testauth"))
			Convey("Then no error should be returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
