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

	"github.com/northwesternmutual/grammes/query/cardinality"
	. "github.com/smartystreets/goconvey/convey"
)

func TestProperty(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Property' is called with object strings", func() {
			result := g.Property("obj1", "obj2", "obj3", "obj4")
			Convey("Then result should equal 'g.property('obj1','obj2','obj3','obj4')", func() {
				So(result.String(), ShouldEqual, "g.property(\"obj1\",\"obj2\",\"obj3\",\"obj4\")")
			})
		})
		Convey("When 'Property' is called with object strings and cardinality", func() {
			result := g.Property(cardinality.Set, "obj1", "obj2")
			Convey("Then result should equal 'g.property(set,'obj1','obj2')'", func() {
				So(result.String(), ShouldEqual, "g.property(set,\"obj1\",\"obj2\")")
			})
		})
		Convey("When 'Property' is called with object strings and ints", func() {
			result := g.Property(3, 4, 4)
			Convey("Then result should equal 'g.property(3,4,4)'", func() {
				So(result.String(), ShouldEqual, "g.property(3,4,4)")
			})
		})

		Convey("When 'Properties' is called with no arguments", func() {
			result := g.Properties()
			Convey("Then result should equal 'g.properties()'", func() {
				So(result.String(), ShouldEqual, "g.properties()")
			})
		})
	})
}
