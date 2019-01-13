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

func TestWhere(t *testing.T) {

	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Where' is called with just first argument string", func() {
			result := g.Where("testFirst")
			Convey("Then result should equal 'g.where(testFirst)'", func() {
				So(result.String(), ShouldEqual, "g.where(testFirst)")
			})
		})

		Convey("When 'Where' is called with just first argument int", func() {
			result := g.Where(1234)
			Convey("Then result should equal 'g.where()'", func() {
				So(result.String(), ShouldEqual, "g.where()")
			})
		})

		Convey("When 'Where' is called with just first argument ) String {", func() {
			result := g.Where(NewTraversal().Label().Raw())
			Convey("Then result should equal 'g.where(label())'", func() {
				So(result.String(), ShouldEqual, "g.where(label())")
			})
		})

		Convey("When 'Where' is called with extras argument multiple string", func() {
			result := g.Where("testFirst", "testExtras1", "testExtras2")
			Convey("Then result should equal 'g.where(testFirst,testExtras1,testExtras2)'", func() {
				So(result.String(), ShouldEqual, "g.where(testFirst,testExtras1,testExtras2)")
			})
		})
	})
}
