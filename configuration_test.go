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
	"strconv"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWithErrorChannel(t *testing.T) {
	Convey("Given an error channel and dialer", t, func() {
		var errs chan error
		dialer := &mockDialer{}
		Convey("When Dial is called with error channel", func() {
			c, _ := Dial(dialer, WithErrorChannel(errs))
			Convey("Then the client error channel should be set", func() {
				So(c.err, ShouldResemble, errs)
			})
		})
	})
}

func TestWithLogger(t *testing.T) {
	Convey("Given a logger and dialer", t, func() {
		var l testLogger
		dialer := &mockDialer{}
		Convey("When Dial is called with logger", func() {
			c, _ := Dial(dialer, WithLogger(l))
			Convey("Then the client logger should be set", func() {
				So(c.logger, ShouldResemble, l)
			})
		})
	})
}

func TestWithGremlinVersion(t *testing.T) {
	Convey("Given a Gremlin version and dialer", t, func() {
		v := 3
		dialer := &mockDialer{}
		Convey("When Dial is called with Gremlin Version", func() {
			c, _ := Dial(dialer, WithGremlinVersion(v))
			Convey("Then the client Gremlin version should be set", func() {
				So(c.gremlinVersion, ShouldEqual, strconv.Itoa(v))
			})
		})
	})
}

func TestWithMaxConcurrentMessages(t *testing.T) {
	Convey("Given an int and dialer", t, func() {
		m := 2
		dialer := &mockDialer{}
		Convey("When Dial is called with max concurrent messages", func() {
			c, _ := Dial(dialer, WithMaxConcurrentMessages(m))
			Convey("Then the client request channel should be set", func() {
				So(c.request, ShouldNotBeNil)
			})
		})
	})
}

func TestWithAuthUserPass(t *testing.T) {
	Convey("Given a username, password and dialer", t, func() {
		user := "testuser"
		pass := "testpass"
		dialer := &mockDialer{}
		Convey("And Dial is called with username and password", func() {
			_, err := Dial(dialer, WithAuthUserPass(user, pass))
			Convey("Then no error should be encountered", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestWithTimeout(t *testing.T) {
	Convey("Given a timeout and dialer", t, func() {
		t := 5 * time.Second
		dialer := &mockDialer{}
		Convey("And Dial is called with timeout", func() {
			_, err := Dial(dialer, WithTimeout(t))
			Convey("Then no error should be encountered", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestWithPingInterval(t *testing.T) {
	Convey("Given a ping interval and dialer", t, func() {
		p := 5 * time.Second
		dialer := &mockDialer{}
		Convey("And Dial is called with ping interval", func() {
			_, err := Dial(dialer, WithPingInterval(p))
			Convey("Then no error should be encountered", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestWithWritingWait(t *testing.T) {
	Convey("Given a writing wait and dialer", t, func() {
		w := 5 * time.Second
		dialer := &mockDialer{}
		Convey("And Dial is called with writing wait", func() {
			_, err := Dial(dialer, WithWritingWait(w))
			Convey("Then no error should be encountered", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestWithReadingWait(t *testing.T) {
	Convey("Given a reading wait and dialer", t, func() {
		r := 5 * time.Second
		dialer := &mockDialer{}
		Convey("And Dial is called with reading wait", func() {
			_, err := Dial(dialer, WithReadingWait(r))
			Convey("Then no error should be encountered", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
