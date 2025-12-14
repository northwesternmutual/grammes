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

	"github.com/northwesternmutual/grammes/query/scope"
)

func TestSkip(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Skip' is called with with just first argument", func() {
			result := g.Skip(scope.Global)
			Convey("Then result should equal 'g.skip(global)'", func() {
				So(result.String(), ShouldEqual, "g.skip(global)")
			})
		})

		Convey("When 'Skip' is called with multiple extraFloat arguments", func() {
			result := g.Skip(scope.Local, 1.234)
			Convey("Then result should equal 'g.skip(local,1.234)'", func() {
				So(result.String(), ShouldEqual, "g.skip(local,1.234f)")
			})
		})
	})
}
