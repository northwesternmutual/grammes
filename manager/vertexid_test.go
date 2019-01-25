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

func TestVertexIDsByString(t *testing.T) {
	Convey("Given a string executor and query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(idResponse), nil }
		qm := newVertexIDQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VertexIDsByString is called", func() {
			_, err := qm.VertexIDsByString("testquery")
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVertexIDsByStringQueryError(t *testing.T) {
	Convey("Given a string executor and query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		qm := newVertexIDQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VertexIDsByString is called and encounters a querying error", func() {
			_, err := qm.VertexIDsByString("testquery")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexIDsByStringUnmarshalError(t *testing.T) {
	defer func() {
		jsonUnmarshal = json.Unmarshal
	}()
	jsonUnmarshal = func([]byte, interface{}) error { return errors.New("ERROR") }
	Convey("Given a string executor and query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(idResponse), nil }
		qm := newVertexIDQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VertexIDsByString is called and encounters an unmarshalling error", func() {
			_, err := qm.VertexIDsByString("testquery")
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexIDByQuery(t *testing.T) {
	Convey("Given a string executor and query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(idResponse), nil }
		qm := newVertexIDQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VertexIDsByQuery is called", func() {
			var q mockQuery
			_, err := qm.VertexIDsByQuery(q)
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVertexIDByQueryError(t *testing.T) {
	Convey("Given a string executor and query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		qm := newVertexIDQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VertexIDsByQuery is called and encounters a querying error", func() {
			var q mockQuery
			_, err := qm.VertexIDsByQuery(q)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexIDs(t *testing.T) {
	Convey("Given a string executor and query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(idResponse), nil }
		qm := newVertexIDQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VertexIDs is called", func() {
			_, err := qm.VertexIDs("testlabel", "prop1", "prop2")
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVertexIDsParamError(t *testing.T) {
	Convey("Given a string executor and query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(idResponse), nil }
		qm := newVertexIDQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VertexIDs is called with an odd number of parameters", func() {
			_, err := qm.VertexIDs("testlabel", "prop1")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexIDsQueryError(t *testing.T) {
	Convey("Given a string executor and query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		qm := newVertexIDQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VertexIDs is called and encounters a querying error", func() {
			_, err := qm.VertexIDs("testlabel", "prop1", "prop2")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
