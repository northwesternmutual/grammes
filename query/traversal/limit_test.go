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

package traversal

import (
	"testing"

	"github.com/northwesternmutual/grammes/query/scope"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLimit(t *testing.T) {
	// Convey("Given a ) String { that represents the graph's traversal", t, func() {
	// 	g := NewTraversal()
	// 	Convey("When 'Limit' is called with a string", func() {
	// 		result := g.Limit("myVertex")
	// 		Convey("Then result should equal 'g.addV('myVertex')'", func() {
	// 			So(result.String(), ShouldEqual, "g.addV('myVertex')")
	// 		})
	// 	})

	// 	Convey("When 'AddV' is called with a traversal", func() {
	// 		result := g.AddV(NewTraversal().Label().Raw())
	// 		Convey("Then result should equal 'g.addV(label())'", func() {
	// 			So(result.String(), ShouldEqual, "g.addV(label())")
	// 		})
	// 	})
	// })
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Limit' is called with an int", func() {
			result := g.Limit(123)
			Convey("Then result should equal 'g.limit(123)'", func() {
				So(result.String(), ShouldEqual, "g.limit(123)")
			})
		})

		Convey("When 'Limit' is called with no params", func() {
			result := g.Limit()
			Convey("Then result should equal 'g.limit(\"", func() {
				So(result.String(), ShouldEqual, "g.limit(")
			})
		})

		Convey("When 'Limit' is called with too many params", func() {
			result := g.Limit(1, 2, 3)
			Convey("Then result should equal 'g.limit(1)2)3)'", func() {
				So(result.String(), ShouldEqual, "g.limit(1)2)3)")
			})
		})

		Convey("When 'Limit' is called with scope", func() {
			result := g.Limit(scope.Local)
			Convey("Then result should equal 'g.limit(local,'", func() {
				So(result.String(), ShouldEqual, "g.limit(local,")
			})
		})
	})
}
