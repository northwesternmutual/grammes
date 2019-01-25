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

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/northwesternmutual/grammes/gremconnect"
)

func TestLaunchConnection(t *testing.T) {
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}

	Convey("Given a client", t, func() {
		dialer := &mockDialerStruct{}
		dialer.connect = nil
		c, _ := mockDial(dialer)
		Convey("and launchConnection() is called", func() {
			err := c.launchConnection()
			Convey("Then the err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestClose(t *testing.T) {
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}

	Convey("Given a client", t, func() {
		dialer := &mockDialerStruct{}
		dialer.connect = nil
		c, _ := Dial(dialer)
		Convey("Then no errors or panics should be thrown when Close() is called", func() {
			c.Close()
		})
	})
}

func TestIsConnected(t *testing.T) {
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
	}()

	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	Convey("Given a client", t, func() {
		dialer := &mockDialerStruct{}
		dialer.connect = nil
		c, _ := Dial(dialer)
		Convey("When IsConnected() is called", func() {
			connection := c.IsConnected()
			Convey("Then the return value should match the isConnected var", func() {
				So(connection, ShouldEqual, dialer.isConnected)
			})
		})
	})
}

func TestRedial(t *testing.T) {
	Convey("Given a client", t, func() {
		dialer := &mockDialerStruct{}
		c, _ := mockDial(dialer)
		Convey("When Redial is called", func() {
			err := c.Redial(dialer)
			Convey("Then the error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestConnect(t *testing.T) {
	tempNewWebSocketDialer := NewWebSocketDialer
	defer func() {
		NewWebSocketDialer = tempNewWebSocketDialer
	}()
	NewWebSocketDialer = func(string) gremconnect.Dialer { return &mockDialerStruct{} }
	Convey("Given a client", t, func() {
		dialer := &mockDialerStruct{}
		c, _ := mockDial(dialer)
		c.conn = dialer
		Convey("And Connect is called with the client", func() {
			dialer.isDisposed = true
			err := c.Connect()
			Convey("Then no error should be returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestConnectNoConnection(t *testing.T) {
	Convey("Given a client", t, func() {
		c := setupClient()
		Convey("When Connect is called without a connection", func() {
			err := c.Connect()
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestConnectErrorLaunchingConnection(t *testing.T) {
	Convey("Given a client", t, func() {
		dialer := &mockDialerStruct{}
		c := setupClient()
		c.conn = dialer
		Convey("When Connect is called and launching the connection throws an error", func() {
			dialer.isDisposed = true
			dialer.connect = errors.New("ERROR")
			err := c.Connect()
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
