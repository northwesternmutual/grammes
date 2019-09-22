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

package manager

import (
	"encoding/json"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnmarshalID(t *testing.T) {
	Convey("Given a byte ID response", t, func() {
		Convey("When unmarshalID is called", func() {
			id, _ := unmarshalID([][]byte{[]byte(vertexResponse)})
			Convey("Then the return ID should be 0", func() {
				So(id, ShouldEqual, 0)
			})
		})
	})
}

func TestUnmarshalIDError(t *testing.T) {
	defer func() {
		jsonUnmarshal = json.Unmarshal
	}()
	jsonUnmarshal = func([]byte, interface{}) error { return errors.New("ERROR") }
	Convey("Given a byte ID response", t, func() {
		Convey("When unmarshalID is called and encounters an error", func() {
			_, err := unmarshalID([][]byte{[]byte(vertexResponse)})
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
