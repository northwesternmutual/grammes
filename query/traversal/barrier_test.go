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

func TestBarrier(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Barrier' is called with an integer", func() {
			result := g.Barrier(4)
			Convey("Then result should equal 'g.barrier(4)'", func() {
				So(result.String(), ShouldEqual, "g.barrier(4)")
			})
		})

		Convey("When 'Barrier' is called with a Consumer string", func() {
			result := g.Barrier("consumer")
			Convey("Then result should equal 'g.barrier(consumer)'", func() {
				So(result.String(), ShouldEqual, "g.barrier(consumer)")
			})
		})

		Convey("When 'Barrier' is called with no parameters", func() {
			result := g.Barrier()
			Convey("Then result should equal 'g.barrier()'", func() {
				So(result.String(), ShouldEqual, "g.barrier()")
			})
		})
		Convey("When 'Barrier' is called with too many parameters", func() {
			result := g.Barrier("param1", "param2")
			Convey("Then result should equal 'g.barrier(param1)'", func() {
				So(result.String(), ShouldEqual, "g.barrier(param1)")
			})
		})
	})
}
