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
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/northwesternmutual/grammes/gremconnect"
	"github.com/northwesternmutual/grammes/gremerror"
)

func TestExecuteRequest(t *testing.T) {

	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}

	Convey("Given a client that represents the Gremlin client", t, func() {
		dialer := &mockDialerStruct{}
		dialer.response = newVertexResponse
		c, _ := Dial(dialer)
		Convey("When 'executeRequest' is called with query", func() {
			q := "testQuery"
			var b, r map[string]string
			res, err := c.executeRequest(q, nil, b, r, nil)
			Convey("Then err should be nil and the test result should be returned", func() {
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
	})
}

func TestWriteWorkerErrorWriting(t *testing.T) {
	t.Parallel()

	Convey("Given a client that represents the Gremlin client", t, func() {
		dialer := &mockDialerWriteError{}
		c, _ := Dial(dialer)
		Convey("When there is an error writing the message", func() {
			errReceived := false
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				for {
					select {
					case <-c.err:
						errReceived = true
						wg.Done()
					default:
						continue
					}
				}
			}()
			c.dispatchRequest([]byte(newVertexResponse))
			wg.Wait()
			Convey("Then the error should be sent through the channel", func() {
				So(errReceived, ShouldBeTrue)
			})
		})
	})
}

func TestExecuteRequestErrorPreparingRequest(t *testing.T) {

	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
		gremPrepareRequest = gremconnect.PrepareRequest
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	gremPrepareRequest = func(string, *time.Duration, map[string]string, map[string]string, *uuid.UUID, map[string]string) (gremconnect.Request, string, error) {
		var req gremconnect.Request
		return req, "test", errors.New("ERROR")
	}
	Convey("Given a client that represents the Gremlin client", t, func() {
		dialer := &mockDialerStruct{}
		c, _ := Dial(dialer)
		Convey("When 'executeRequest' is called and preparing the request throws an error", func() {
			bindings := make(map[string]string)
			rebindings := make(map[string]string)
			_, err := c.executeRequest("testing", nil, bindings, rebindings, nil)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestExecuteRequestErrorPackagingRequest(t *testing.T) {

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
		dialer := &mockDialerStruct{}
		c, _ := Dial(dialer)
		Convey("When 'executeRequest' is called and packaging the request throws an error", func() {
			bindings := make(map[string]string)
			rebindings := make(map[string]string)
			_, err := c.executeRequest("testing", nil, bindings, rebindings, nil)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestExecuteRequestErrorRetrievingResponse(t *testing.T) {

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
		dialer := &mockDialerStruct{}
		dialer.response = `
			{
				"requestId": "61616161-6161-6161-2d61-6161612d6161",
				"status": {
					"message": "",
					"code": 597,
					"attributes": {}
				},
				"result":{"data":null,"meta":{"@type":"g:Map","@value":[]}}
			}
			`
		c, _ := Dial(dialer)
		Convey("When 'executeRequest' is called and retrieving the response throws an error", func() {
			bindings := make(map[string]string)
			rebindings := make(map[string]string)
			_, err := c.executeRequest("testing", nil, bindings, rebindings, nil)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("Then the error should be gremerror.NetworkError", func() {
				_, ok := err.(*gremerror.NetworkError)
				So(ok, ShouldBeTrue)
			})
		})
	})
}

func TestAuthenticate(t *testing.T) {

	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	Convey("Given a client that represents the Gremlin client", t, func() {
		dialer := &mockDialerStruct{}
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

	defer func() {
		gremconnect.GenUUID = uuid.NewUUID
	}()
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return uuid.UUID(a), nil
	}
	Convey("Given a client without auth credentials", t, func() {
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
		dialer := &mockDialerStruct{}
		dialer.response = newVertexResponse
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
		dialer := &mockDialerStruct{}
		dialer.response = newVertexResponse
		c, _ := Dial(dialer)
		Convey("When authenticate is called and packaging the request throws an error", func() {
			err := c.authenticate("testauth")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestExecuteRequestErrorTimeout(t *testing.T) {
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return a, nil
	}
	Convey("Given a client that represents the Gremlin client", t, func() {
		dialer := &mockDialerStruct{}
		dialer.readDuration = 10 * time.Minute
		c, _ := Dial(dialer)
		c.requestTimeout = time.Second
		Convey("When 'executeRequest' is called and timeout occurs", func() {
			bindings := make(map[string]string)
			rebindings := make(map[string]string)
			_, err := c.executeRequest("testing", nil, bindings, rebindings, nil)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("Then the error should be timeout", func() {
				So(err.Error(), ShouldEqual, "request failed with timeout")
			})
		})
	})
}

func TestExecuteRequestErrorQueryTimeout(t *testing.T) {
	gremconnect.GenUUID = func() (uuid.UUID, error) {
		var a [16]byte
		copy(a[:], "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
		return a, nil
	}
	Convey("Given a client that represents the Gremlin client", t, func() {
		dialer := &mockDialerStruct{}
		dialer.readDuration = 10 * time.Minute
		c, _ := Dial(dialer)
		c.requestTimeout = 10 * time.Minute
		queryTimeout := time.Second
		Convey("When 'executeRequest' is called and timeout occurs", func() {
			bindings := make(map[string]string)
			rebindings := make(map[string]string)
			_, err := c.executeRequest("testing", &queryTimeout, bindings, rebindings, nil)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("Then the error should be timeout", func() {
				So(err.Error(), ShouldEqual, "request failed with timeout")
			})
		})
	})
}
