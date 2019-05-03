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

func TestVerticesByString(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VerticesByString is called", func() {
			_, err := qm.VerticesByString("testquery")
			Convey("Then the return error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVerticesByStringQueryError(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VerticesByString is called and encounters an error", func() {
			_, err := qm.VerticesByString("testquery")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVerticesByStringUnmarshalError(t *testing.T) {
	defer func() {
		jsonUnmarshal = json.Unmarshal
	}()
	jsonUnmarshal = func([]byte, interface{}) error { return errors.New("ERROR") }
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VerticesByString is called and there is an error unmarshalling", func() {
			_, err := qm.VerticesByString("testquery")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVerticesByQuery(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VerticesByString is called", func() {
			var q mockQuery
			_, err := qm.VerticesByQuery(q)
			Convey("Then the return error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVerticesByQueryError(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VerticesByString is called and encounters an error", func() {
			var q mockQuery
			_, err := qm.VerticesByQuery(q)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAllVertices(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When AllVertices is called", func() {
			_, err := qm.AllVertices()
			Convey("Then the return error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAllVerticesError(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When AllVertices is called and encounters an error", func() {
			_, err := qm.AllVertices()
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertexByID(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VertexByID is called", func() {
			_, err := qm.VertexByID(1234)
			Convey("Then the return error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVertexByIDError(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When VertexByID is called and encounters an error", func() {
			_, err := qm.VertexByID(1234)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVertices(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When Vertices is called", func() {
			_, err := qm.Vertices("testlabel", "prop1", "prop2")
			Convey("Then the return error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestVerticesPropertyError(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When Vertices is called with an odd number of properties", func() {
			_, err := qm.Vertices("testlabel", "prop1")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestVerticesQueryError(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		qm := newGetVertexQueryManager(logging.NewBasicLogger(), execute)
		Convey("When Vertices is called and encounters a querying error", func() {
			_, err := qm.Vertices("testlabel", "prop1", "prop2")
			Convey("Then the return error should be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
