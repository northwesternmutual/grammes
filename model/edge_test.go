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

package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPropertyValue(t *testing.T) {
	Convey("Given a variable that represents the Edge struct", t, func() {
		e := Edge{Type: "tesType"}

		Convey("When 'PropertyValue' is called with a string", func() {
			result := e.PropertyValue("testKey")
			Convey("Then result should equal nil", func() {
				So(result, ShouldEqual, nil)
			})
		})
	})
}

func TestID(t *testing.T) {
	Convey("Given a variable that represents the Edge struct", t, func() {
		// erid := EdgeRelationID{RelationID: "testRelID"}
		// ev := EdgeValue{ID: EdgeID{Value: erid}}
		e := Edge{Type: "tesType", Value: EdgeValue{ID: EdgeID{Value: EdgeRelationID{RelationID: "testRelID"}}}}

		Convey("When 'ID' is called", func() {
			result := e.ID()
			Convey("Then result should be 'testRelID'", func() {
				So(result, ShouldEqual, "testRelID")
			})
		})
	})
}

func TestLabel(t *testing.T) {
	Convey("Given a variable that represents the Edge struct", t, func() {
		ev := EdgeValue{Label: "testLabel"}
		e := Edge{Type: "tesType", Value: ev}

		Convey("When 'Label' is called", func() {
			result := e.Label()
			Convey("Then result should be 'testLabel'", func() {
				So(result, ShouldEqual, "testLabel")
			})
		})
	})
}

func TestOutVertexID(t *testing.T) {
	Convey("Given a variable that represents the Edge struct", t, func() {
		// erid := EdgeRelationID{RelationID: "testRelID"}
		// ev := EdgeValue{ID: EdgeID{Value: erid}}
		e := Edge{Type: "tesType", Value: EdgeValue{OutV: EdgeVertex{Value: 12345}}}

		Convey("When 'OutVertexID' is called", func() {
			result := e.OutVertexID()
			Convey("Then result should be 12345", func() {
				So(result, ShouldEqual, 12345)
			})
		})
	})
}

func TestInVertexID(t *testing.T) {
	Convey("Given a variable that represents the Edge struct", t, func() {
		e := Edge{Type: "tesType", Value: EdgeValue{InV: EdgeVertex{Value: 54321}}}

		Convey("When 'InVertexID' is called", func() {
			result := e.InVertexID()
			Convey("Then result should be 54321", func() {
				So(result, ShouldEqual, 54321)
			})
		})
	})
}

func TestOutVertexLabel(t *testing.T) {
	Convey("Given a variable that represents the Edge struct", t, func() {

		e := Edge{Type: "tesType", Value: EdgeValue{OutVLabel: "testOutLabel"}}

		Convey("When 'GetOutVertexLabel' is called", func() {
			result := e.OutVertexLabel()
			Convey("Then result should be 'testOutLabel'", func() {
				So(result, ShouldEqual, "testOutLabel")
			})
		})
	})
}

func TestInVertexLabel(t *testing.T) {
	Convey("Given a variable that represents the Edge struct", t, func() {

		e := Edge{Type: "tesType", Value: EdgeValue{InVLabel: "testInLabel"}}

		Convey("When 'InVertexLabel' is called", func() {
			result := e.InVertexLabel()
			Convey("Then result should be 'testInLabel'", func() {
				So(result, ShouldEqual, "testInLabel")
			})
		})
	})
}

func TestQueryOutVertex(t *testing.T) {
	Convey("Given a variable that represents the Edge struct", t, func() {
		e := Edge{Type: "tesType", Value: EdgeValue{OutV: EdgeVertex{Value: 12345}}}

		Convey("When 'QueryOutVertex' is called with a string", func() {
			var client queryClient
			result, _ := e.QueryOutVertex(client)
			var nilVert Vertex
			Convey("Then result should equal nil", func() {
				So(result, ShouldResemble, nilVert)
			})
		})
	})
}

func TestQueryInVertex(t *testing.T) {
	Convey("Given a variable that represents the Edge struct", t, func() {
		e := Edge{Type: "tesType", Value: EdgeValue{OutV: EdgeVertex{Value: 12345}}}

		Convey("When 'QueryInVertex' is called with a string", func() {
			var client queryClient
			result, _ := e.QueryInVertex(client)
			var nilVert Vertex
			Convey("Then result should equal nil", func() {
				So(result, ShouldResemble, nilVert)
			})
		})
	})
}
