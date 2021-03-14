package googleinterview_test

import (
	gi "golang-coding-problems/internal/googleinterview"

	"testing"

	"github.com/franela/goblin"
)

func TestMoveZerosToLeft(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Move Zeros to the left", func() {
		arr, expected := []int{1, 0, 5, 9, 0, 3, 20, 0}, []int{0, 0, 0, 1, 5, 9, 3, 20}
		g.It("Should move all zeroes to the left", func() {
			g.Assert(gi.MoveZerosToLeft(arr)).Eql(expected)
		})
	})
}
