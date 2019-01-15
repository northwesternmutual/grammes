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

func TestMax(t *testing.T) {
	Convey("Given a string that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Max' is called with no arguments", func() {
			result := g.Max()
			Convey("Then result should equal 'g.max()'", func() {
				So(result.String(), ShouldEqual, "g.max()")
			})
		})

		Convey("When 'Max' is called with one Scope", func() {
			result := g.Max(scope.Global)
			Convey("Then result should equal 'g.max(global)'", func() {
				So(result.String(), ShouldEqual, "g.max(global)")
			})
		})

		Convey("When 'Max' is called with multiple scopes", func() {
			result := g.Max(scope.Global, scope.Local)
			Convey("Then result should equal 'g.max()'", func() {
				So(result.String(), ShouldEqual, "g.max()")
			})
		})
	})
}
