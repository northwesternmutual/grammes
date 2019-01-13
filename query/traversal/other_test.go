package traversal

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateGUUID(t *testing.T) {
	Convey("Given a String that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'GenerateGUUID' is called", func() {
			result := g.OtherV()
			Convey("Then result should equal 'g.otherV()'", func() {
				So(result.String(), ShouldEqual, "g.otherV()")
			})
		})
	})
}
