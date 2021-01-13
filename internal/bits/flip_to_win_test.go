package bits_test

import (
	"golang-coding-problems/internal/bits"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestFlipToWin(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Flip to Win", func() {

		g.It("Should return the max consecutive ones if there we flip one bit", func() {
			g.Assert(bits.FlipToWin(1775)).Eql(8)
			g.Assert(bits.FlipToWin(205)).Eql(4)
			g.Assert(bits.FlipToWin(217)).Eql(5)
		})

	})
}
