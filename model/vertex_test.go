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

func TestVertexPropertyValue(t *testing.T) {
	Convey("Given a variable that represents the Vertex struct", t, func() {
		testID := PropertyID{Value: "relID"}
		testValue := ValueWrapper{PropertyDetailedValue: PropertyDetailedValue{Value: "tstInterface", Type: "pdvType"}}
		pi := Property{Type: "piType", Value: PropertyValue{ID: testID, Value: testValue, Label: "tstLabel"}}

		pdmap := map[string][]Property{"testKey": {pi}}

		v := Vertex{Type: "tesType", Value: VertexValue{Properties: pdmap}}

		Convey("When 'PropertyValue' is called with a string and int", func() {
			result := v.PropertyValue("testKey", 0)
			Convey("Then result should equal 'tstInterface'", func() {
				So(result, ShouldEqual, "tstInterface")
			})
		})
	})
}

func TestVertexID(t *testing.T) {
	Convey("Given a variable that represents the Vertex struct", t, func() {
		testID := PropertyID{Value: "relID"}
		testValue := ValueWrapper{PropertyDetailedValue: PropertyDetailedValue{Value: "tstInterface", Type: "pdvType"}}
		pi := Property{Type: "piType", Value: PropertyValue{ID: testID, Value: testValue, Label: "tstLabel"}}
		pdmap := map[string][]Property{"testKey": {pi}}

		v := Vertex{Type: "tesType", Value: VertexValue{Properties: pdmap, ID: 6789}}

		Convey("When 'ID' is called with a string and int", func() {
			result := v.ID()
			Convey("Then result should equal 6789", func() {
				So(result, ShouldEqual, 6789)
			})
		})
	})
}

func TestVertexLabel(t *testing.T) {
	Convey("Given a variable that represents the Vertex struct", t, func() {
		testID := PropertyID{Value: "relID"}
		testValue := ValueWrapper{PropertyDetailedValue: PropertyDetailedValue{Value: "tstInterface", Type: "pdvType"}}
		pi := Property{Type: "piType", Value: PropertyValue{ID: testID, Value: testValue, Label: "tstLabel"}}
		pdmap := map[string][]Property{"testKey": {pi}}

		v := Vertex{Type: "tesType", Value: VertexValue{Label: "testLabel", ID: 6789, Properties: pdmap}}

		Convey("When 'Label' is called with a string and int", func() {
			result := v.Label()
			Convey("Then result should equal 'testLabel'", func() {
				So(result, ShouldEqual, "testLabel")
			})
		})
	})
}

func TestQueryBothEdges(t *testing.T) {
	Convey("Given a variable that represents the Vertex struct", t, func() {

		v := Vertex{Type: "tesType"}
		Convey("When 'QueryBothEdges' is called with a string and int", func() {
			var client queryClient
			var edges []Edge
			result, _ := v.QueryBothEdges(client)
			Convey("Then result should equal []Edge", func() {
				So(result, ShouldEqual, edges)
			})
		})
	})
}

func TestQueryOutEdges(t *testing.T) {
	Convey("Given a variable that represents the Vertex struct", t, func() {

		v := Vertex{Type: "tesType"}
		Convey("When 'QueryOutEdges' is called with a string and int", func() {
			var client queryClient
			var edges []Edge
			result, _ := v.QueryOutEdges(client)
			Convey("Then result should equal []Edge", func() {
				So(result, ShouldEqual, edges)
			})
		})
	})
}

func TestQueryInEdges(t *testing.T) {
	Convey("Given a variable that represents the Vertex struct", t, func() {

		v := Vertex{Type: "tesType"}
		Convey("When 'QueryInEdges' is called with a string and int", func() {
			var client queryClient
			var edges []Edge
			result, _ := v.QueryInEdges(client)
			Convey("Then result should equal []Edge", func() {
				So(result, ShouldEqual, edges)
			})
		})
	})
}

func TestAddEdge(t *testing.T) {
	Convey("Given a variable that represents the Vertex struct", t, func() {

		v := Vertex{Type: "tesType"}
		Convey("When 'AddEdges' is called with a string and int", func() {
			var client queryClient
			var edge Edge
			var tstoutVID int64
			tstoutVID = 12345

			result, _ := v.AddEdge(client, "testLbl", tstoutVID, "tstIntrf1", "tstIntrf2", 7777, 9876)
			Convey("Then result should equal 'testLabel'", func() {
				So(result, ShouldResemble, edge)
			})
		})
	})
}
