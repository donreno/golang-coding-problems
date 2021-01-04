package arrays_test

import (
	"testing"

	goblin "github.com/franela/goblin"

	a "golang-coding-problems/internal/arrays"
)

func TestIsRotation(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Check if string B is rotation of A", func() {

		g.It("Should identify that 'erbottlewat' is rotation of 'waterbottle'", func() {
			g.Assert(a.IsRotation("waterbottle", "erbottlewat")).IsTrue()
		})

		g.It("Should identify that 'armadillo' is not a rotation of 'beer'", func() {
			g.Assert(a.IsRotation("beer", "armadillo")).IsFalse()
		})
	})
}
