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

package grammes

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/northwesternmutual/grammes/query/graph"
	"github.com/northwesternmutual/grammes/query/traversal"
)

func TestTraversal(t *testing.T) {
	t.Parallel()

	Convey("Given we call the Traversal function", t, func() {
		expected := traversal.NewTraversal()
		t := Traversal()
		Convey("Then the return value should match the expected result", func() {
			So(t, ShouldResemble, expected)
		})
	})
}

func TestCustomTraversal(t *testing.T) {
	t.Parallel()

	Convey("Given a query string", t, func() {
		q := "testQuery"
		expected := traversal.NewCustomTraversal(q)
		Convey("When CustomTraversal is called", func() {
			t := CustomTraversal(q)
			Convey("Then the return value should match the expected result", func() {
				So(t, ShouldResemble, expected)
			})
		})
	})
}

func TestVerboseTraversal(t *testing.T) {
	t.Parallel()

	Convey("Given we call the VerboseTraversal function", t, func() {
		expected := graph.String("graph")
		t := VerboseTraversal()
		Convey("Then the return value should match the expected result", func() {
			So(t, ShouldResemble, expected)
		})
	})
}
