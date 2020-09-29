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

func TestAddStep(t *testing.T) {
	Convey("Given a graph traversal", t, func() {
		g := NewTraversal()
		Convey("When AddStep is called with []byte", func() {
			b := []byte("test")
			g.AddStep("test", b)
			Convey("Then g should equal g.test(test)", func() {
				So(g.String(), ShouldEqual, "g.test(test)")
			})
		})

		Convey("When AddStep is called with byte", func() {
			b := byte('T')
			g.AddStep("test", b)
			Convey("Then g should equal g.test(T)", func() {
				So(g.String(), ShouldEqual, "g.test(T)")
			})
		})

		Convey("When AddStep is called with int32", func() {
			i := int32(1234)
			g.AddStep("test", i)
			Convey("Then g should equal g.test(1234)", func() {
				So(g.String(), ShouldEqual, "g.test(1234)")
			})
		})

		Convey("When AddStep is called with int64", func() {
			i := int64(1234)
			g.AddStep("test", i)
			Convey("Then g should equal g.test(1234)", func() {
				So(g.String(), ShouldEqual, "g.test(1234)")
			})
		})

		Convey("When AddStep is called with float64", func() {
			i := float64(1.234)
			g.AddStep("test", i)
			Convey("Then g should equal g.test(1.234)", func() {
				So(g.String(), ShouldEqual, "g.test(1.234f)")
			})
		})

		Convey("When AddStep is called with bool", func() {
			b := true
			g.AddStep("test", b)
			Convey("Then g should equal g.test(true)", func() {
				So(g.String(), ShouldEqual, "g.test(true)")
			})
		})

		Convey("When AddStep should escape characters", func() {
			s := `a 'test' with \ and \\ and nested {"key": "value", "key2": "val\"val"} and {'key': 'value', 'kay2': 'val\'val'}`
			g.AddStep("test", s)
			expected := `g.test('a \'test\' with \\ and \\\\ and nested {"key": "value", "key2": "val\\"val"} and {\'key\': \'value\', \'kay2\': \'val\\\'val\'}')`
			Convey("Then g should equal "+expected, func() {
				So(g.String(), ShouldEqual, expected)
			})
		})
	})

}

func TestString(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'String' is called", func() {
			result := g.String()
			Convey("Then result should equal 'g'", func() {
				So(result, ShouldEqual, "g")
			})
		})
	})
}

func TestRaw(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Raw' is called", func() {
			result := g.Raw()
			Convey("Then result should equal 'g'", func() {
				So(result.String(), ShouldEqual, "g")
			})
		})
	})
}

func TestGatherInts(t *testing.T) {
	Convey("When GatherInts is called with one param", t, func() {
		result := gatherInts(1)
		Convey("Then result should equal '1'", func() {
			So(result, ShouldEqual, "1")
		})
	})

	Convey("When GatherInts is called with many params", t, func() {
		result := gatherInts(1, 2, 3)
		Convey("Then result should equal empty string", func() {
			So(result, ShouldEqual, "")
		})
	})

	Convey("When GatherInts is called with no params", t, func() {
		result := gatherInts()
		Convey("Then result should equal empty string", func() {
			So(result, ShouldEqual, "")
		})
	})
}
