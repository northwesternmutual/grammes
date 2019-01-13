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

	. "github.com/smartystreets/goconvey/convey"
)

func TestPageRank(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'PageRank' is called with no arguments", func() {
			result := g.PageRank()
			Convey("Then result should equal 'g.pageRank()'", func() {
				So(result.String(), ShouldEqual, "g.pageRank()")
			})
		})

		Convey("When 'PageRank' is called with no multiple arguments", func() {
			var flt1, flt2 float32
			flt1 = 1.234
			flt2 = 5.678
			result := g.PageRank(flt1, flt2)
			Convey("Then result should equal 'g.pageRank(1.234000)'", func() {
				So(result.String(), ShouldEqual, "g.pageRank(1.234000)")
			})
		})
		// Convey("When 'PageRank' is called with a multiple Scope", func() {
		// 	scopeL := scope.Local
		// 	scopeG := scope.Global
		// 	result := g.PageRank(scopeL, scopeG)
		// 	Convey("Then result should equal 'g.min(global)'", func() {
		// 		So(result.String(), ShouldEqual, "g.min(local)")
		// 	})
		// })
	})
}
