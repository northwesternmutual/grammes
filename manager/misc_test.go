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

	"github.com/northwesternmutual/grammes/logging"
)

func TestDropAll(t *testing.T) {
	Convey("Given a string executor and misc query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, nil }
		mm := newMiscQueryManager(logging.NewBasicLogger(), execute)
		Convey("When DropAll is called", func() {
			err := mm.DropAll()
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestSetVertexProperty(t *testing.T) {
	Convey("Given a string executor and misc query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, nil }
		mm := newMiscQueryManager(logging.NewBasicLogger(), execute)
		Convey("When SetVertexProperty is called", func() {
			err := mm.SetVertexProperty(1234, "prop1", "prop2")
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestSetVertexPropertyParameterError(t *testing.T) {
	Convey("Given a string executor and misc query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, nil }
		mm := newMiscQueryManager(logging.NewBasicLogger(), execute)
		Convey("When SetVertexProperty is called with an odd number of properties", func() {
			err := mm.SetVertexProperty(1234, "prop1")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestSetVertexPropertyQueryError(t *testing.T) {
	Convey("Given a string executor and misc query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		mm := newMiscQueryManager(logging.NewBasicLogger(), execute)
		Convey("When SetVertexProperty is called and encounters a querying error", func() {
			err := mm.SetVertexProperty(1234)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexCount(t *testing.T) {
	Convey("Given a string executor and misc query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(idResponse), nil }
		mm := newMiscQueryManager(logging.NewBasicLogger(), execute)
		Convey("When SetVertexProperty is called", func() {
			c, _ := mm.VertexCount()
			Convey("Then the count should equal 255", func() {
				So(c, ShouldEqual, 255)
			})
		})
	})
}

func TestVertexCountQueryError(t *testing.T) {
	Convey("Given a string executor and misc query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		mm := newMiscQueryManager(logging.NewBasicLogger(), execute)
		Convey("When SetVertexProperty is called and encounters a querying error", func() {
			_, err := mm.VertexCount()
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexCountUnmarshalError(t *testing.T) {
	defer func() {
		JSONUnmarshal = json.Unmarshal
	}()
	JSONUnmarshal = func([]byte, interface{}) error { return errors.New("ERROR") }
	Convey("Given a string executor and misc query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, nil }
		mm := newMiscQueryManager(logging.NewBasicLogger(), execute)
		Convey("When SetVertexProperty is called and encounters an numarshalling error", func() {
			_, err := mm.VertexCount()
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
