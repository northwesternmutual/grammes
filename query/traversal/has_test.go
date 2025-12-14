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
	"go/token"
	"testing"

	"github.com/northwesternmutual/grammes/query/predicate"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHas(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Has' is called with object strings", func() {
			result := g.Has("obj1", "obj2", "obj3")
			Convey("Then result should equal 'g.has('obj1','obj2','obj3')'", func() {
				So(result.String(), ShouldEqual, "g.has('obj1','obj2','obj3')")
			})
		})

		Convey("When 'Has' is called with a traversal", func() {
			result := g.Has("testHas", NewTraversal().Label().Raw())
			Convey("Then result should equal 'g.has('testHas',label())'", func() {
				So(result.String(), ShouldEqual, "g.has('testHas',label())")
			})
		})

		Convey("When 'Has' is called with one token parameter", func() {
			var t token.Token
			result := g.Has(t)
			Convey("Then result should equal 'g.has(ILLEGAL)'", func() {
				So(result.String(), ShouldEqual, "g.has(ILLEGAL)")
			})
		})

		Convey("When 'Has' is called with one int", func() {
			result := g.Has(1234)
			Convey("Then result should equal 'g.has(1234)'", func() {
				So(result.String(), ShouldEqual, "g.has(1234)")
			})
		})

		Convey("When 'Has' is called with too many params", func() {
			result := g.Has("first", "second", "third", "fourth")
			Convey("Then result should equal 'g.has('first','second','third','fourth')'", func() {
				So(result.String(), ShouldEqual, "g.has('first','second','third','fourth')")
			})
		})

		Convey("When 'Has' is called with many different param types", func() {
			p := new(predicate.Predicate)
			*p = "predicate"
			result := g.Has("first", p, 1234)
			Convey("Then result should equal 'g.has('first',predicate,1234')'", func() {
				So(result.String(), ShouldEqual, "g.has('first',predicate,1234)")
			})
		})
	})
}

func TestHasID(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'HasID' is called with one parameter", func() {
			result := g.HasID("tstObjOrP")
			Convey("Then result should equal 'g.hasId('tstObjOrP')'", func() {
				So(result.String(), ShouldEqual, "g.hasId('tstObjOrP')")
			})
		})

		Convey("When 'HasID' is called with a multiple params", func() {
			result := g.HasID("tstObjOrP", "tstObj1", "tstObj2")
			Convey("Then result should equal 'g.hasId('tstObjOrP','tstObj1','tstObj2')'", func() {
				So(result.String(), ShouldEqual, "g.hasId('tstObjOrP','tstObj1','tstObj2')")
			})
		})
	})
}

func TestHasKey(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'HasKey' is called with one parameter", func() {
			result := g.HasKey("tstpOrStr")
			Convey("Then result should equal 'g.hasKey('tstpOrStr')'", func() {
				So(result.String(), ShouldEqual, "g.hasKey('tstpOrStr')")
			})
		})

		Convey("When 'HasKey' is called with one int parameter", func() {
			result := g.HasKey(1234)
			Convey("Then result should equal 'g.hasKey(1234)'", func() {
				So(result.String(), ShouldEqual, "g.hasKey(1234)")
			})
		})

		Convey("When 'HasKey' is called with a multiple params", func() {
			result := g.HasKey("tstpOrStr", "tstHandled1", "tstHandled2")
			Convey("Then result should equal 'g.hasKey('tstpOrStr','tstHandled1','tstHandled2')'", func() {
				So(result.String(), ShouldEqual, "g.hasKey('tstpOrStr','tstHandled1','tstHandled2')")
			})
		})
	})
}

func TestHasLabel(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'HasLabel' is called with one parameter", func() {
			result := g.HasLabel("tstpOrStr")
			Convey("Then result should equal 'g.hasLabel('tstpOrStr')'", func() {
				So(result.String(), ShouldEqual, "g.hasLabel('tstpOrStr')")
			})
		})

		Convey("When 'HasLabel' is called with a multiple params", func() {
			result := g.HasLabel("tstpOrStr", "tstHandled1", "tstHandled2")
			Convey("Then result should equal 'g.hasLabel('tstpOrStr','tstHandled1','tstHandled2')'", func() {
				So(result.String(), ShouldEqual, "g.hasLabel('tstpOrStr','tstHandled1','tstHandled2')")
			})
		})

		Convey("When 'HasLabel' is called with a predicate", func() {
			result := g.HasLabel(predicate.LessThan(12))
			Convey("Then result should equal 'g.hasLabel(lt(12))'", func() {
				So(result.String(), ShouldEqual, "g.hasLabel(lt(12))")
			})
		})

	})
}

func TestHasNot(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'HasNot' is called", func() {
			result := g.HasNot("testStr")
			Convey("Then result should equal 'g.hasNot('testStr')'", func() {
				So(result.String(), ShouldEqual, "g.hasNot('testStr')")
			})
		})
	})
}

func TestHasValue(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'HasValue' is called with one parameter", func() {
			result := g.HasValue("tstObjOrP")
			Convey("Then result should equal 'g.hasValue('tstObjOrP')'", func() {
				So(result.String(), ShouldEqual, "g.hasValue('tstObjOrP')")
			})
		})

		Convey("When 'HasValue' is called with one int parameter", func() {
			result := g.HasValue(1234)
			Convey("Then result should equal 'g.hasValue(1234)'", func() {
				So(result.String(), ShouldEqual, "g.hasValue(1234)")
			})
		})

		Convey("When 'HasValue' is called with a multiple params", func() {
			result := g.HasValue("tstObjOrP", "tstObj1", "tstObj2")
			Convey("Then result should equal 'g.hasValue('tstObjOrP','tstObj1','tstObj2')'", func() {
				So(result.String(), ShouldEqual, "g.hasValue('tstObjOrP','tstObj1','tstObj2')")
			})
		})
	})
}
