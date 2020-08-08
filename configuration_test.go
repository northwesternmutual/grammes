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
	"net/http"
	"strconv"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWithErrorChannel(t *testing.T) {
	t.Parallel()

	Convey("Given an error channel and dialer", t, func() {
		var errs chan error
		dialer := &mockDialerStruct{}
		Convey("When Dial is called with error channel", func() {
			c, _ := Dial(dialer, WithErrorChannel(errs))
			Convey("Then the client error channel should be set", func() {
				So(c.err, ShouldResemble, errs)
			})
		})
	})
}

func TestWithLogger(t *testing.T) {
	t.Parallel()

	Convey("Given a logger and dialer", t, func() {
		dialer := &mockDialerStruct{}
		Convey("When Dial is called with logger", func() {
			c, _ := mockDial(dialer, WithLogger(dialer.logger))
			Convey("Then the client logger should be set", func() {
				So(c.logger, ShouldResemble, dialer.logger)
			})
		})
	})
}

func TestWithGremlinVersion(t *testing.T) {
	t.Parallel()

	Convey("Given a Gremlin version and dialer", t, func() {
		v := 3
		dialer := &mockDialerStruct{}
		Convey("When Dial is called with Gremlin Version", func() {
			c, _ := mockDial(dialer, WithGremlinVersion(v))
			Convey("Then the client Gremlin version should be set", func() {
				So(c.gremlinVersion, ShouldEqual, strconv.Itoa(v))
			})
		})
	})
}

func TestWithMaxConcurrentMessages(t *testing.T) {
	t.Parallel()

	Convey("Given an int and dialer", t, func() {
		m := 2
		dialer := &mockDialerStruct{}
		Convey("When Dial is called with max concurrent messages", func() {
			c, _ := mockDial(dialer, WithMaxConcurrentMessages(m))
			Convey("Then the client request channel should be set", func() {
				So(c.request, ShouldNotBeNil)
			})
		})
	})
}

func TestWithAuthUserPass(t *testing.T) {
	t.Parallel()

	Convey("Given a username, password and dialer", t, func() {
		user := "testuser"
		pass := "testpass"
		dialer := &mockDialerStruct{}
		Convey("And Dial is called with username and password", func() {
			_, err := mockDial(dialer, WithAuthUserPass(user, pass))
			Convey("Then no error should be encountered", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestWithHTTPAuth(t *testing.T) {
	t.Parallel()

	Convey("Given an auth provider and dialer", t, func() {
		dialer := &mockDialerStruct{}
		Convey("And Dial is called with username and password", func() {
			_, err := mockDial(dialer, WithHTTPAuth(func(request *http.Request) error {
				return nil
			}))
			Convey("Then no error should be encountered", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestWithTimeout(t *testing.T) {
	t.Parallel()

	Convey("Given a timeout and dialer", t, func() {
		dialer := &mockDialerStruct{}
		dialer.timeout = 5 * time.Second
		Convey("And Dial is called with timeout", func() {
			_, err := mockDial(dialer, WithTimeout(dialer.timeout))
			Convey("Then no error should be encountered", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestWithPingInterval(t *testing.T) {
	t.Parallel()

	Convey("Given a ping interval and dialer", t, func() {
		dialer := &mockDialerStruct{}
		dialer.pingInterval = 5 * time.Second
		Convey("And Dial is called with ping interval", func() {
			_, err := mockDial(dialer, WithPingInterval(dialer.pingInterval))
			Convey("Then no error should be encountered", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestWithWritingWait(t *testing.T) {
	t.Parallel()

	Convey("Given a writing wait and dialer", t, func() {
		dialer := &mockDialerStruct{}
		dialer.writingWait = 5 * time.Second
		Convey("And Dial is called with writing wait", func() {
			_, err := mockDial(dialer, WithWritingWait(dialer.writingWait))
			Convey("Then no error should be encountered", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestWithReadingWait(t *testing.T) {
	t.Parallel()

	Convey("Given a reading wait and dialer", t, func() {
		dialer := &mockDialerStruct{}
		dialer.readingWait = 5 * time.Second
		Convey("And Dial is called with reading wait", func() {
			_, err := mockDial(dialer, WithReadingWait(dialer.readingWait))
			Convey("Then no error should be encountered", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
