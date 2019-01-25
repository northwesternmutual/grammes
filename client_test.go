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

func TestDial(t *testing.T) {
	Convey("Given a mock dialer", t, func() {
		dialer := &mockDialerStruct{}
		Convey("And connection is successful", func() {
			Convey("When Dial is called with the mock dialer", func() {
				c, _ := Dial(dialer, WithLogger(dialer.logger))
				Convey("Then c.Schema should not be nil", func() {
					So(c.IsBroken(), ShouldBeFalse)
				})
			})

		})

		Convey("And connection is unsuccessful", func() {
			dialer.connect = errors.New("ERR")
			Convey("When Dial is called with the mock dialer", func() {
				c, _ := Dial(dialer, WithLogger(dialer.logger))
				Convey("Then c.Schema should not be nil", func() {
					So(c.IsBroken(), ShouldBeTrue)
				})
			})

		})
	})
}

func TestDialWithWebSocket(t *testing.T) {
	tempNewWebSocketDialer := NewWebSocketDialer
	defer func() {
		NewWebSocketDialer = tempNewWebSocketDialer
		gremconnect.GenUUID = uuid.NewUUID
	}()
	NewWebSocketDialer = func(string) gremconnect.Dialer { return &mockDialerStruct{} }
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	Convey("Given a host string and error channel", t, func() {
		host := ""
		Convey("And connection is successful", func() {
			Convey("When NewClient is called", func() {
				c, _ := DialWithWebSocket(host)
				Convey("Then client shouldn't be nil", func() {
					So(c, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestSetLogger(t *testing.T) {
	Convey("Given a client and a logger", t, func() {
		dialer := &mockDialerStruct{}
		c, _ := Dial(dialer)
		Convey("and SetLogger is called", func() {
			c.SetLogger(dialer.logger)
			Convey("Then the client logger should resemble l", func() {
				So(c.logger, ShouldResemble, dialer.logger)
			})
		})
	})
}

func TestIsBroken(t *testing.T) {
	Convey("Given a client", t, func() {
		dialer := &mockDialerStruct{}
		c, _ := Dial(dialer)
		c.broken = false
		Convey("And IsBroken is called", func() {
			b := c.IsBroken()
			Convey("Then the return should equal client.broken", func() {
				So(b, ShouldEqual, c.broken)
			})
		})
	})
}

func TestAddress(t *testing.T) {
	Convey("Given a client", t, func() {
		dialer := &mockDialerStruct{}
		c, _ := Dial(dialer)
		Convey("And Address is called", func() {
			a := c.Address()
			Convey("Then the return should equal the mock connection address", func() {
				So(a, ShouldEqual, "")
			})
		})
	})
}

func TestAuth(t *testing.T) {
	Convey("Given a client", t, func() {
		dialer := &mockDialerStruct{}
		c, _ := Dial(dialer)
		Convey("And Auth is called", func() {
			auth, _ := c.Auth()
			Convey("Then the returned auth should not be nil", func() {
				So(auth, ShouldNotBeNil)
			})
		})
	})
}
