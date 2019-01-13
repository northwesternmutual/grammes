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
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/northwesternmutual/grammes/gremconnect"
)

type mockDialerAuthError gremconnect.WebSocket

func (*mockDialerAuthError) Connect() error     { return connect }
func (*mockDialerAuthError) Close() error       { return nil }
func (*mockDialerAuthError) Write([]byte) error { return nil }
func (m *mockDialerAuthError) Read() ([]byte, error) {
	if readCount < 1 {
		time.Sleep(100 * time.Millisecond)
		readCount++
		m.Quit <- struct{}{}
		return []byte(response), nil
	}
	return nil, nil
}
func (*mockDialerAuthError) Ping(chan error)   {}
func (*mockDialerAuthError) IsConnected() bool { return isConnected }
func (*mockDialerAuthError) IsDisposed() bool  { return isDisposed }
func (*mockDialerAuthError) Auth() (*gremconnect.Auth, error) {
	return &gremconnect.Auth{}, errors.New("ERROR")
}
func (*mockDialerAuthError) Address() string { return "" }
func (m *mockDialerAuthError) GetQuit() chan struct{} {
	m.Quit = make(chan struct{})
	return m.Quit
}
func (*mockDialerAuthError) SetAuth(string, string) {}
func (*mockDialerAuthError) SetTimeout(int)         {}
func (*mockDialerAuthError) SetPingInterval(int)    {}
func (*mockDialerAuthError) SetWritingWait(int)     {}
func (*mockDialerAuthError) SetReadingWait(int)     {}

func TestExecuteRequest(t *testing.T) {
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
		dialer := &mockDialer{}
		c, _ := Dial(dialer)
		Convey("When 'executeRequest' is called with query", func() {
			bindings := make(map[string]string)
			rebindings := make(map[string]string)
			readCount = 0
			result, err := c.executeRequest("testing", bindings, rebindings)
			Convey("Then err should be nil and the test result should be returned", func() {
				So(err, ShouldBeNil)
				So(result, ShouldNotBeNil)
			})
		})
	})
}

func TestExecuteRequestErrorPreparingRequest(t *testing.T) {
	readCount = 1
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
		gremPrepareRequest = gremconnect.PrepareRequest
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	gremPrepareRequest = func(string, map[string]string, map[string]string) (gremconnect.Request, string, error) {
		var req gremconnect.Request
		return req, "test", errors.New("ERROR")
	}
	Convey("Given a client that represents the Gremlin client", t, func() {
		connect = nil
		dialer := &mockDialer{}
		c, _ := Dial(dialer)
		Convey("When 'executeRequest' is called and preparing the request throws an error", func() {
			bindings := make(map[string]string)
			rebindings := make(map[string]string)
			readCount = 0
			_, err := c.executeRequest("testing", bindings, rebindings)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestExecuteRequestErrorPackagingRequest(t *testing.T) {
	readCount = 1
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
		gremPackageRequest = gremconnect.PackageRequest
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	gremPackageRequest = func(gremconnect.Request, string) ([]byte, error) { return nil, errors.New("ERROR") }
	Convey("Given a client that represents the Gremlin client", t, func() {
		connect = nil
		dialer := &mockDialer{}
		c, _ := Dial(dialer)
		Convey("When 'executeRequest' is called and packaging the request throws an error", func() {
			bindings := make(map[string]string)
			rebindings := make(map[string]string)
			readCount = 0
			_, err := c.executeRequest("testing", bindings, rebindings)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestExecuteRequestErrorRetrievingResponse(t *testing.T) {
	readCount = 1
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
		jsonMarshalData = json.Marshal
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	jsonMarshalData = func(interface{}) ([]byte, error) { return nil, errors.New("ERROR") }
	Convey("Given a client that represents the Gremlin client", t, func() {
		connect = nil
		dialer := &mockDialer{}
		c, _ := Dial(dialer)
		Convey("When 'executeRequest' is called and retrieving the response throws an error", func() {
			bindings := make(map[string]string)
			rebindings := make(map[string]string)
			readCount = 0
			_, err := c.executeRequest("testing", bindings, rebindings)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

// func TestWriteWorkerErrorWriting(t *testing.T) {
// 	readCount = 0
// 	defer func() {
// 		gremconnect.GenUUID = uuid.NewUUID
// 	}()
// 	gremconnect.GenUUID = func() (uuid.UUID, error) {
// 	var a [16]byte
// 	copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
// 	return uuid.UUID(a), nil
// }
// 	Convey("Given a client", t, func() {
// 		connect = nil
// 		dialer := &mockDialerWriteError{}
// 		var l TestLogger
// 		c, _ := Dial(dialer, WithLogger(l))
// 		Convey("When writeWorker is executed but returns an error", func() {
// 			errCounter := 0
// 			for range c.err {
// 				errCounter++
// 			}
// 			close(c.err)
// 			Convey("Then the error should be send down the error channel", func() {
// 				So(errCounter, ShouldNotEqual, 0)
// 			})
// 		})
// 	})
// }

func TestAuthenticate(t *testing.T) {
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
		dialer := &mockDialer{}
		c, _ := Dial(dialer)
		Convey("When 'authenticate' is called with query", func() {
			err := c.authenticate("requestIDtest")
			Convey("Then result should be empty", func() {
				So(err, ShouldEqual, nil)
			})
		})
	})
}

func TestAuthenticateAuthError(t *testing.T) {
	readCount = 1
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	Convey("Given a client without auth credentials", t, func() {
		connect = nil
		dialer := &mockDialerAuthError{}
		c, _ := Dial(dialer)
		Convey("When authenticate is called", func() {
			err := c.authenticate("testauth")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAuthenticateErrorPraparingAuthRequest(t *testing.T) {
	readCount = 1
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
		gremPrepareAuthRequest = gremconnect.PrepareAuthRequest
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	gremPrepareAuthRequest = func(string, string, string) (gremconnect.Request, error) {
		return gremconnect.Request{}, errors.New("ERROR")
	}
	Convey("Given a client", t, func() {
		connect = nil
		dialer := &mockDialer{}
		c, _ := Dial(dialer)
		Convey("When authenticate is called and preparing the request throws an error", func() {
			err := c.authenticate("testauth")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAuthenticateErrorPackagingRequest(t *testing.T) {
	readCount = 1
	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
		gremPackageRequest = gremconnect.PackageRequest
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	gremPackageRequest = func(gremconnect.Request, string) ([]byte, error) { return nil, errors.New("ERROR") }
	Convey("Given a client that represents the Gremlin client", t, func() {
		connect = nil
		dialer := &mockDialer{}
		c, _ := Dial(dialer)
		Convey("When authenticate is called and packaging the request throws an error", func() {
			err := c.authenticate("testauth")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
