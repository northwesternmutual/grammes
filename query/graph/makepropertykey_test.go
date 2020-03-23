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

package graph

import (
	"testing"

	"github.com/northwesternmutual/grammes/query/cardinality"
	"github.com/northwesternmutual/grammes/query/datatype"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMakePropertyKey(t *testing.T) {
	Convey("Given a *String that represents the verbose graph traversal", t, func() {
		g := NewGraph()
		Convey("When 'MakePropertyKey' is called with a label string, DataType and Cardinality", func() {

			var dt = datatype.String
			var cr = cardinality.List
			result := g.MakePropertyKey("labelTest", dt, cr)
			Convey("Then result should equal 'graph.makePropertyKey('labelTest').dataType(String.class).cardinality(list)'", func() {
				So(result.String(), ShouldEqual, "graph.makePropertyKey(\"labelTest\").dataType(String.class).cardinality(list)")
			})
		})
	})
}
