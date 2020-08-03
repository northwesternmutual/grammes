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
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
)

var upgrader = websocket.Upgrader{}

var echo = func(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		err = c.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func TestConnect(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(echo))
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	defer s.Close()

	Convey("Given a dialer object", t, func() {
		dialer := &WebSocket{}
		dialer.address = u

		Convey("And we call Connect() using the dialer", func() {
			err := dialer.Connect()

			Convey("Then the err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("and dialer.Connected should be true", func() {
				So(dialer.connected, ShouldEqual, true)
			})
		})
		// Convey("And when we test the pong handler", func() {
		// 	// dialer.conn.PongHandler()
		// 	// pongErr := handler("testmsg")
		// 	// fmt.Println(pongErr)
		// 	// pongErr2 := dialer.conn.handler("testmsg")
		// 	// fmt.Println(pongErr2)
		// })
	})
}

func TestIsConnected(t *testing.T) {
	Convey("Given a WebSocket", t, func() {
		dialer := &WebSocket{}

		Convey("And we set 'Connected' to true", func() {
			dialer.connected = true

			Convey("Then the IsConnected function should return true", func() {
				conn := dialer.IsConnected()
				So(conn, ShouldEqual, dialer.connected)
			})
		})
	})

}

func TestIsDisposed(t *testing.T) {
	Convey("Given a WebSocket", t, func() {
		dialer := &WebSocket{}

		Convey("And we set Disposed to true", func() {
			dialer.disposed = true

			Convey("Then the IsDisposed function should return true", func() {
				disp := dialer.IsDisposed()
				So(disp, ShouldEqual, dialer.disposed)
			})
		})
	})
}

func TestWriteAndRead(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(echo))
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	defer s.Close()
	Convey("Given a Websocket connection to a test server", t, func() {
		dialer := &WebSocket{}
		dialer.address = u
		_ = dialer.Connect()
		defer dialer.conn.Close()

		Convey("And a message is written and read", func() {
			writemsg := []byte("test")
			writeerr := dialer.Write(writemsg)
			readmsg, readerr := dialer.Read()

			Convey("Then the err should be nil and we should retrieve the written message from the server", func() {
				So(writeerr, ShouldBeNil)
				So(readerr, ShouldBeNil)
				So(readmsg, ShouldResemble, writemsg)
			})
		})
	})
}

func TestClose(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(echo))
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	defer s.Close()
	Convey("Given a WebSocket connection to a test server", t, func() {
		dialer := &WebSocket{}
		dialer.address = u
		dialer.Quit = make(chan struct{})
		_ = dialer.Connect()

		Convey("And we close the connection", func() {
			err := dialer.Close()

			Convey("Then err should be nil and dialer.Disposed should be true", func() {
				So(err, ShouldBeNil)
				So(dialer.disposed, ShouldBeTrue)
			})
		})
	})
}

func TestAuthValid(t *testing.T) {
	Convey("Given a WebSocket with auth credentials", t, func() {
		dialer := &WebSocket{}
		testAuth := &Auth{Username: "testuser", Password: "testpassword"}
		dialer.auth = testAuth

		Convey("And we make a call to the GetAuth() function", func() {
			auth, _ := dialer.Auth()

			Convey("Then the return value should be the credentials we just assigned", func() {
				So(auth, ShouldResemble, dialer.auth)
			})
		})
	})
}

func TestAuthInvalid(t *testing.T) {
	Convey("Given a WebSocket without auth credentials", t, func() {
		dialer := &WebSocket{}
		auth, err := dialer.Auth()
		Convey("Then the GetAuth() function should throw an error", func() {
			So(err, ShouldNotBeNil)
		})
		Convey("Then auth should equal nil", func() {
			So(auth, ShouldBeNil)
		})
	})
}

func TestAddress(t *testing.T) {
	Convey("Given a WebSocket with an address", t, func() {
		dialer := &WebSocket{}
		dialer.address = "testaddress"

		Convey("And the GetAddress() function is called", func() {
			address := dialer.Address()

			Convey("Then the return value should match the assigned address", func() {
				So(address, ShouldEqual, dialer.address)
			})
		})
	})
}
func TestGetQuit(t *testing.T) {
	Convey("Given a WebSocket with a quit channel", t, func() {
		dialer := &WebSocket{}
		dialer.Quit = make(chan struct{})

		Convey("And the GetQuit() function is called", func() {
			q := dialer.GetQuit()

			Convey("Then the return value should match the assigned quit channel", func() {
				So(q, ShouldResemble, dialer.Quit)
			})
		})
	})
}

func TestPing(t *testing.T) {
	Convey("Given a WebSocket connection with a ping interval and writing wait time", t, func() {
		dialer := &WebSocket{}
		dialer.pingInterval = 75 * time.Millisecond
		dialer.writingWait = 10 * time.Millisecond
		dialer.Quit = make(chan struct{})

		s := httptest.NewServer(http.HandlerFunc(echo))
		u := "ws" + strings.TrimPrefix(s.URL, "http")
		defer s.Close()

		dialer.address = u
		dialer.Connect()

		Convey("And we ping the test server", func() {
			errs := make(chan error)
			go dialer.Ping(errs)
			time.Sleep(500 * time.Millisecond)
			dialer.Close()
			close(errs)

			Convey("Then we should not receive any errors in the return channel", func() {
				errCounter := 0
				for range errs {
					errCounter++
				}
				So(errCounter, ShouldEqual, 0)
			})
		})
	})
}

func TestPingLostConnection(t *testing.T) {
	Convey("Given a WebSocket connection with a ping interval and writing wait time", t, func() {
		dialer := &WebSocket{}
		dialer.pingInterval = 75 * time.Millisecond
		dialer.writingWait = 0 * time.Millisecond
		dialer.Quit = make(chan struct{})

		s := httptest.NewServer(http.HandlerFunc(echo))
		u := "ws" + strings.TrimPrefix(s.URL, "http")
		defer s.Close()

		dialer.address = u
		dialer.Connect()

		Convey("And we call Ping without allowing enough time for a response", func() {
			errs := make(chan error, 1)
			go dialer.Ping(errs)
			time.Sleep(100 * time.Millisecond)
			dialer.Close()
			close(errs)
			Convey("Then we should receive an error in the return channel", func() {
				errCounter := 0
				for range errs {
					errCounter++
				}
				So(errCounter, ShouldEqual, 1)
			})
		})
	})
}
func TestSetAuth(t *testing.T) {
	Convey("Given a WebSocket, username and password", t, func() {
		dialer := &WebSocket{}
		user := "testuser"
		pass := "testpass"
		Convey("And the GetQuit() function is called", func() {
			dialer.SetAuth(user, pass)
			Convey("Then the username and password for the dialer should be set", func() {
				So(dialer.auth.Username, ShouldEqual, "testuser")
				So(dialer.auth.Password, ShouldEqual, "testpass")
			})
		})
	})
}

func TestSetTimeout(t *testing.T) {
	Convey("Given a WebSocket, and timeout interval", t, func() {
		dialer := &WebSocket{}
		t := 5 * time.Second
		Convey("And SetTimeout is called", func() {
			dialer.SetTimeout(t)
			Convey("Then the timeout should be set in the dialer", func() {
				So(dialer.timeout, ShouldEqual, t)
			})
		})
	})
}

func TestSetPingInterval(t *testing.T) {
	Convey("Given a WebSocket and a ping interval", t, func() {
		dialer := &WebSocket{}
		p := 5 * time.Second
		Convey("And SetPingInterval is called", func() {
			dialer.SetPingInterval(p)
			Convey("Then the ping interval should be set in the dialer", func() {
				So(dialer.pingInterval, ShouldEqual, p)
			})
		})
	})
}

func TestSetWritingWait(t *testing.T) {
	Convey("Given a WebSocket and a writing wait", t, func() {
		dialer := &WebSocket{}
		w := 5 * time.Second
		Convey("And SetWritingWait is called", func() {
			dialer.SetWritingWait(w)
			Convey("Then the writing wait should be set in the dialer", func() {
				So(dialer.writingWait, ShouldEqual, w)
			})
		})
	})
}

func TestSetReadingWait(t *testing.T) {
	Convey("Given a WebSocket and a reading wait", t, func() {
		dialer := &WebSocket{}
		r := 5 * time.Second
		Convey("And SetReadingWait is called", func() {
			dialer.SetReadingWait(r)
			Convey("Then the reading wait should be set in the dialer", func() {
				So(dialer.readingWait, ShouldEqual, r)
			})
		})
	})
}

func TestSetWriteBufferSize(t *testing.T) {
	Convey("Given a WebSocket and a write buffer size", t, func() {
		dialer := &WebSocket{}
		writeBufferSize := 512 * 1024
		Convey("And SetWriteBufferSize is called", func() {
			dialer.SetWriteBufferSize(writeBufferSize)
			Convey("Then the write buffer size should be set in the dialer", func() {
				So(dialer.writeBufferSize, ShouldEqual, writeBufferSize)
			})
		})
	})
}

func TestSetReadBufferSize(t *testing.T) {
	Convey("Given a WebSocket and a read buffer size", t, func() {
		dialer := &WebSocket{}
		readBufferSize := 256 * 1024
		Convey("And SetReadBufferSize is called", func() {
			dialer.SetReadBufferSize(readBufferSize)
			Convey("Then the read buffer size should be set in the dialer", func() {
				So(dialer.readBufferSize, ShouldEqual, readBufferSize)
			})
		})
	})
}

func TestHandshakeTimeout(t *testing.T) {
	Convey("Given a WebSocket and a handshake timeout", t, func() {
		dialer := &WebSocket{}
		handshakeTimeout := time.Second
		Convey("And SetHandshakeTimeout is called", func() {
			dialer.SetHandshakeTimeout(handshakeTimeout)
			Convey("Then the handshake timeout should be set in the dialer", func() {
				So(dialer.handshakeTimeout, ShouldEqual, handshakeTimeout)
			})
		})
	})
}

func TestCompression(t *testing.T) {
	Convey("Given a WebSocket and a compression flag", t, func() {
		dialer := &WebSocket{}
		enableCompression := true
		Convey("And SetCompression is called", func() {
			dialer.SetCompression(enableCompression)
			Convey("Then the compression flag should be set in the dialer", func() {
				So(dialer.enableCompression, ShouldEqual, enableCompression)
			})
		})
	})
}
