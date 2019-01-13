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

	"github.com/northwesternmutual/grammes/query/direction"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTo(t *testing.T) {

	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'To' is called with just first argument string", func() {
			result := g.To("testFirst")
			Convey("Then result should equal 'g.to(testFirst)'", func() {
				So(result.String(), ShouldEqual, "g.to(testFirst)")
			})
		})

		Convey("When 'To' is called with just first argument int", func() {
			result := g.To(1234)
			Convey("Then result should equal 'g.to()'", func() {
				So(result.String(), ShouldEqual, "g.to()")
			})
		})

		Convey("When 'To' is called with just first argument ) String {", func() {
			result := g.To(NewTraversal().Label())
			Convey("Then result should equal 'g.to(label())'", func() {
				So(result.String(), ShouldEqual, "g.to(label())")
			})
		})

		Convey("When 'To' is called with just first argument pop.Pop", func() {
			var p direction.Direction
			p = "testDirection"
			result := g.To(p)
			Convey("Then result should equal 'g.to(testDirection)'", func() {
				So(result.String(), ShouldEqual, "g.to(testDirection)")
			})
		})

		Convey("When 'To' is called with extras argument multiple string", func() {
			result := g.To("testFirst", "testExtras1", "testExtras2")
			Convey("Then result should equal 'g.to(testFirst,testExtras1,testExtras2)'", func() {
				So(result.String(), ShouldEqual, "g.to(testFirst,testExtras1,testExtras2)")
			})
		})
	})
}

func TestToE(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'ToE' is called", func() {
			var dir direction.Direction
			dir = "left"
			result := g.ToE(dir, "testStr")
			Convey("Then result should equal 'g.toE(left, 'testStr')'", func() {
				So(result.String(), ShouldEqual, "g.toE(left, 'testStr')")
			})
		})
	})
}

func TestToV(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'ToV' is called", func() {
			var dir direction.Direction
			dir = "left"
			result := g.ToV(dir)
			Convey("Then result should equal 'g.toV(left)'", func() {
				So(result.String(), ShouldEqual, "g.toV(left)")
			})
		})
	})
}

func TestToVId(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'ToVId' is called", func() {
			result := g.ToVId(1234)
			Convey("Then result should equal 'g.to(V().hasId(1234))'", func() {
				So(result.String(), ShouldEqual, "g.to(V().hasId(1234))")
			})
		})
	})
}
