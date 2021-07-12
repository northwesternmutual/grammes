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
	"encoding/json"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrepareRequest(t *testing.T) {
	Convey("Given a query string, binding, and rebindings", t, func() {
		query := ""
		bindings := make(map[string]string)
		rebindings := make(map[string]string)

		Convey("And a request is prepared", func() {
			req, id, err := PrepareRequest(query, nil, bindings, rebindings, nil)

			Convey("Then the request and id should not be nil", func() {
				So(req, ShouldNotBeNil)
				So(id, ShouldNotBeNil)
			})

			Convey("And the error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestPackageRequest(t *testing.T) {
	Convey("Given a request", t, func() {
		req := Request{
			RequestID: "testID",
			Op:        "eval",
			Processor: "",
			Args:      make(map[string]interface{}),
		}

		Convey("And the request is packaged", func() {
			msg, err := PackageRequest(req, "3")

			Convey("Then msg should not be nil", func() {
				So(msg, ShouldNotBeNil)
			})

			Convey("And the error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestPackageRequestJsonMarshalError(t *testing.T) {
	defer func() {
		jsonMarshal = json.Marshal
	}()
	testErr := errors.New("ERROR")
	jsonMarshal = func(interface{}) ([]byte, error) { return nil, testErr }
	Convey("Given a request", t, func() {
		var req Request
		Convey("And PackageRequest throws an error", func() {
			msg, err := PackageRequest(req, "3")
			Convey("Then msg should be nil", func() {
				So(msg, ShouldBeNil)
			})
			Convey("Then err should say 'ERROR'", func() {
				So(err, ShouldEqual, testErr)
			})
		})
	})
}

func TestPrepareAuthRequest(t *testing.T) {
	Convey("Given a request ID, username, and password", t, func() {
		id := "testID"
		user := "testUser"
		pass := "testPass"

		Convey("And and auth request is prepared", func() {
			req, err := PrepareAuthRequest(id, user, pass)

			Convey("Then the request should not be nil", func() {
				So(req, ShouldNotBeNil)
			})

			Convey("And the error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
