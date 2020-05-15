package traversal

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBoth(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Both' is called with an multiple parametrs", func() {
			result := g.Both("lblTest1", "lblTest2")
			Convey("Then result should equal 'g.both('lblTest1','lblTest2')'", func() {
				So(result.String(), ShouldEqual, "g.both('lblTest1','lblTest2')")
			})
		})
	})
}

func TestBothE(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'BothE' is called with an multiple parametrs", func() {
			result := g.BothE("lblTest1", "lblTest2")
			Convey("Then result should equal 'g.bothE('lblTest1','lblTest2')'", func() {
				So(result.String(), ShouldEqual, "g.bothE('lblTest1','lblTest2')")
			})
		})
	})
}
func TestBothV(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'BothV' is called", func() {
			result := g.BothV()
			Convey("Then result should equal 'g.bothV()'", func() {
				So(result.String(), ShouldEqual, "g.bothV()")
			})
		})
	})
}
