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

	. "github.com/smartystreets/goconvey/convey"

	T "github.com/northwesternmutual/grammes/query/token"
)

func TestAddVertex(t *testing.T) {
	Convey("Given a *String that represents the verbose graph traversal", t, func() {
		g := NewGraph()
		Convey("When 'AddVertex' is called with a Token and Label string", func() {
			result := g.AddVertex(T.Label, "testinglabel")
			Convey("Then result should equal 'graph.addVertex(T.label,'testinglabel')'", func() {
				So(result.String(), ShouldEqual, "graph.addVertex(T.label,'testinglabel')")
			})
		})
		Convey("When 'AddVertex' is called with a string and a Token", func() {
			result := g.AddVertex("teststring", "testval", T.Key)
			Convey("Then result should equal 'graph.addVertex('teststring','testval',T.key)'", func() {
				So(result.String(), ShouldEqual, "graph.addVertex('teststring','testval',T.key)")
			})
		})
		Convey("When 'AddVertex' is called with a Token, Label string, "+
			"and property strings", func() {
			result := g.AddVertex(T.Label, "testinglabel", "testkey", "testval")
			Convey("Then result should equal "+
				"'graph.addVertex(T.label,'testinglabel','testkey','testval')'", func() {
				So(result.String(), ShouldEqual, "graph.addVertex(T.label,'testinglabel','testkey','testval')")
			})
		})
	})
}
