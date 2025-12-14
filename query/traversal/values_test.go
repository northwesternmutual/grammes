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

func TestValues(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()

		Convey("When 'Values' is called with no arguments", func() {
			result := g.Values()
			Convey("Then result should equal 'g.values()'", func() {
				So(result.String(), ShouldEqual, "g.values()")
			})
		})

		Convey("When 'Values' is called with a single argument", func() {
			result := g.Values("test")
			Convey("Then result should equal 'g.values('test')'", func() {
				So(result.String(), ShouldEqual, "g.values('test')")
			})
		})

		Convey("When 'Values' is called with a multiple arguments", func() {
			result := g.Values("test1", "test2", "test3")
			Convey("Then result should equal 'g.values('myVertex')'", func() {
				So(result.String(), ShouldEqual, "g.values('test1','test2','test3')")
			})
		})
	})
}
