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
	"github.com/northwesternmutual/grammes/model"
)

var testVertex = model.Vertex{
	Type: "testType",
	Value: model.VertexValue{
		ID:    1234,
		Label: "testLabel",
		Properties: model.PropertyMap{"testDetail": []model.Property{
			{
				Type: "testType",
				Value: model.PropertyValue{
					ID: model.PropertyID{
						Type:  "testIDType",
						Value: "testRelID",
					},
					Value: model.ValueWrapper{
						PropertyDetailedValue: model.PropertyDetailedValue{
							Value: 1234,
							Type:  "testType",
						},
						Partial: true,
					},
					Label: "testLabel",
				},
			}},
		},
	},
}

type mockQuery string

func (mockQuery) String() string { return "TEST" }

func TestAddAPIVertex(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([][]byte, error) { return [][]byte{[]byte(vertexResponse)}, nil }
		qm := newAddVertexQueryManager(logging.NewNilLogger(), execute)
		Convey("When AddAPIVertex is called", func() {
			var data model.APIData
			data.Properties = map[string]string{"testkey": "testval"}
			v, _ := qm.AddAPIVertex(data)
			Convey("Then the return vertex ID value should be 28720", func() {
				So(v.Value.ID, ShouldEqual, 28720)
			})
		})
	})
}

func TestAddAPIVertexError(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([][]byte, error) { return nil, errors.New("ERROR") }
		qm := newAddVertexQueryManager(logging.NewNilLogger(), execute)
		Convey("When AddAPIVertex is called and an error occurs", func() {
			var data model.APIData
			_, err := qm.AddAPIVertex(data)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexByStruct(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([][]byte, error) { return [][]byte{[]byte(vertexResponse)}, nil }
		qm := newAddVertexQueryManager(logging.NewNilLogger(), execute)
		Convey("When AddVertexByStruct is called", func() {
			res, _ := qm.AddVertexByStruct(testVertex)
			Convey("Then the return vertex ID value should be 28720", func() {
				So(res.Value.ID, ShouldEqual, 28720)
			})
		})
	})
}

func TestAddVertexByStructError(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([][]byte, error) { return nil, errors.New("ERROR") }
		qm := newAddVertexQueryManager(logging.NewNilLogger(), execute)
		Convey("When AddVertexByStruct is called and an error is thrown", func() {
			_, err := qm.AddVertexByStruct(testVertex)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexError(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([][]byte, error) { return [][]byte{[]byte(vertexResponse)}, nil }
		qm := newAddVertexQueryManager(logging.NewNilLogger(), execute)
		Convey("When AddVertex is called with an odd number of parameters", func() {
			_, err := qm.AddVertex("testLabel", "prop1")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexLabels(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([][]byte, error) { return [][]byte{[]byte(vertexResponse)}, nil }
		qm := newAddVertexQueryManager(logging.NewNilLogger(), execute)
		Convey("When AddVertexLabels is called", func() {
			_, err := qm.AddVertexLabels("testlabel")
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddVertexLabelsQueryError(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([][]byte, error) { return nil, errors.New("ERROR") }
		qm := newAddVertexQueryManager(logging.NewNilLogger(), execute)
		Convey("When AddVertexLabels is called and encounters a querying error", func() {
			_, err := qm.AddVertexLabels("testlabel")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexByQuery(t *testing.T) {
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([][]byte, error) { return [][]byte{[]byte(vertexResponse)}, nil }
		qm := newAddVertexQueryManager(logging.NewNilLogger(), execute)
		Convey("When AddVertexByQuery is called", func() {
			var q mockQuery
			res, _ := qm.AddVertexByQuery(q)
			Convey("Then the return vertex ID value should be 28720", func() {
				So(res.Value.ID, ShouldEqual, 28720)
			})
		})
	})
}

func TestAddVertexByStringJsonUnmarshalError(t *testing.T) {
	defer func() {
		jsonUnmarshal = json.Unmarshal
	}()
	jsonUnmarshal = func([]byte, interface{}) error { return errors.New("ERROR") }
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([][]byte, error) { return [][]byte{[]byte(vertexResponse)}, nil }
		qm := newAddVertexQueryManager(logging.NewNilLogger(), execute)
		Convey("When AddVertexByString throws an error while unmarshalling", func() {
			_, err := qm.AddVertexByString("testquery")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddVertexByStringReturnnilVertex(t *testing.T) {
	defer func() {
		jsonUnmarshal = json.Unmarshal
	}()
	jsonUnmarshal = func([]byte, interface{}) error { return nil }
	Convey("Given a string executor and vertex query manager", t, func() {
		execute := func(string) ([][]byte, error) { return [][]byte{[]byte(vertexResponse)}, nil }
		qm := newAddVertexQueryManager(logging.NewNilLogger(), execute)
		Convey("When AddVertexByString is called and no vertices are added", func() {
			res, _ := qm.AddVertexByString("testquery")
			Convey("Then the return value should be the nil vertex", func() {
				So(res, ShouldResemble, nilVertex)
			})
		})
	})
}
