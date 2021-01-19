package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestRobotInAGrid(t *testing.T) {
	g := goblin.Goblin(t)

	gridOK := [][]int{
		{0, 0, 0, 1},
		{0, 1, 0, 1},
		{1, 0, 0, 1},
		{0, 0, 0, 0},
	}

	gridNOK := [][]int{
		{0, 0, 0, 1},
		{0, 1, 0, 1},
		{1, 0, 0, 1},
		{0, 0, 1, 0},
	}

	bigGrid := [][]int{
		{0, 0, 0, 1},
		{0, 1, 0, 1},
		{1, 0, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}

	g.Describe("Robot in a grid", func() {

		g.It("Should find exit to grid", func() {
			g.Assert(r.RobotInAGrid(gridOK)).IsTrue()
		})

		g.It("Should find exit to big grid", func() {
			g.Assert(r.RobotInAGrid(bigGrid)).IsTrue()
		})

		g.It("Should not find exit to grid", func() {
			g.Assert(r.RobotInAGrid(gridNOK)).IsFalse()
		})
	})
}
