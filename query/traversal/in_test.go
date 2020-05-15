package traversal

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIn(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'In' is called with an multiple parametrs", func() {
			result := g.In("lblTest1", "lblTest2")
			Convey("Then result should equal 'g.in('lblTest1','lblTest2')'", func() {
				So(result.String(), ShouldEqual, "g.in('lblTest1','lblTest2')")
			})
		})
	})
}

func TestInE(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'InE' is called with an multiple parametrs", func() {
			result := g.InE("lblTest1", "lblTest2")
			Convey("Then result should equal 'g.inE('lblTest1','lblTest2')'", func() {
				So(result.String(), ShouldEqual, "g.inE('lblTest1','lblTest2')")
			})
		})
	})
}
func TestInV(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'InV' is called", func() {
			result := g.InV()
			Convey("Then result should equal 'g.inV()'", func() {
				So(result.String(), ShouldEqual, "g.inV()")
			})
		})
	})
}
