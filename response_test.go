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
