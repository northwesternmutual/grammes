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
	"github.com/northwesternmutual/grammes/query/cardinality"
	"github.com/northwesternmutual/grammes/query/datatype"
	"github.com/northwesternmutual/grammes/query/multiplicity"
)

func TestAddEdgeLabel(t *testing.T) {
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When AddEdgeLabel is called", func() {
			var m = multiplicity.Simple
			_, err := sm.AddEdgeLabel(m, "testlabel")
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddEdgeLabelQueryError(t *testing.T) {
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When AddEdgeLabel is called and encounters a querying error", func() {
			var m = multiplicity.Simple
			_, err := sm.AddEdgeLabel(m, "testlabel")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddEdgeLabelUnmarshalError(t *testing.T) {
	defer func() {
		jsonUnmarshal = json.Unmarshal
	}()
	jsonUnmarshal = func([]byte, interface{}) error { return errors.New("ERROR") }
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When AddEdgeLabel is called and encounters an unmarshalling error", func() {
			var m = multiplicity.Simple
			_, err := sm.AddEdgeLabel(m, "testlabel")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddEdgeLabels(t *testing.T) {
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When AddEdgeLabels is called", func() {
			var m = multiplicity.Simple
			_, err := sm.AddEdgeLabels(m, "testlabel")
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddEdgeLabelsLabelError(t *testing.T) {
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When AddEdgeLabels is called and encounters a querying error", func() {
			var m = multiplicity.Simple
			_, err := sm.AddEdgeLabels(m, "testlabel1", "testlabel2")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddEdgeLabelsInvalidMultiplicity(t *testing.T) {
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When AddEdgeLabels is called with an invalid multiplicity", func() {
			var m = "BADMULT"
			_, err := sm.AddEdgeLabels(m, "testlabel1")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddEdgeLabelsInvalidLabel(t *testing.T) {
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When AddEdgeLabels is called with an invalid label", func() {
			var m = multiplicity.Simple
			_, err := sm.AddEdgeLabels(m, 1234)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddEdgeLabelsQueryingError(t *testing.T) {
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When AddEdgeLabels is called and encounters a querying error", func() {
			var m = multiplicity.Simple
			_, err := sm.AddEdgeLabels(m, "testlabel")
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddPropertyKey(t *testing.T) {
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When AddPropertyKey is called", func() {
			var d = datatype.String
			var c = cardinality.Single
			_, err := sm.AddPropertyKey("testprop", d, c)
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestAddPropertyKeyQueryError(t *testing.T) {
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When AddPropertyKey is called and encounters a querying error", func() {
			var d = datatype.String
			var c = cardinality.Single
			_, err := sm.AddPropertyKey("testprop", d, c)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestAddPropertyKeyUnmarshalError(t *testing.T) {
	defer func() {
		jsonUnmarshal = json.Unmarshal
	}()
	jsonUnmarshal = func([]byte, interface{}) error { return errors.New("ERROR") }
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When AddPropertyKey is called and encounters an unmarshalling error", func() {
			var d = datatype.String
			var c = cardinality.Single
			_, err := sm.AddPropertyKey("testprop", d, c)
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestCommitSchema(t *testing.T) {
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return []byte(vertexResponse), nil }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When CommmitSchema is called", func() {
			_, err := sm.CommitSchema()
			Convey("Then the error returned should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestCommitSchemaQueryError(t *testing.T) {
	Convey("Given a string executor and schema manager", t, func() {
		execute := func(string) ([]byte, error) { return nil, errors.New("ERROR") }
		sm := newSchemaManager(logging.NewBasicLogger(), execute)
		Convey("When CommmitSchema is called and encounters a querying error", func() {
			_, err := sm.CommitSchema()
			Convey("Then the error should be returned", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
