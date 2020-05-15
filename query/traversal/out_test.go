package traversal

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOut(t *testing.T) {
	Convey("Given a String that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Out' is called with an multiple parametrs", func() {
			result := g.Out("lblTest1", "lblTest2")
			Convey("Then result should equal 'g.out('lblTest1','lblTest2')'", func() {
				So(result.String(), ShouldEqual, "g.out('lblTest1','lblTest2')")
			})
		})
	})
}

func TestOutE(t *testing.T) {
	Convey("Given a String that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'OutE' is called with an multiple parametrs", func() {
			result := g.OutE("lblTest1", "lblTest2")
			Convey("Then result should equal 'g.outE('lblTest1','lblTest2')'", func() {
				So(result.String(), ShouldEqual, "g.outE('lblTest1','lblTest2')")
			})
		})
	})
}
func TestOutV(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'OutV' is called", func() {
			result := g.OutV()
			Convey("Then result should equal 'g.outV()'", func() {
				So(result.String(), ShouldEqual, "g.outV()")
			})
		})
	})
}
