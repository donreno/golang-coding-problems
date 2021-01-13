package bits_test

import (
	"golang-coding-problems/internal/bits"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestBitSwapsRequired(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Bit Swaps required to convert numbers A to B", func() {

		g.It("Should return the right number of swaps to convert number", func() {
			g.Assert(bits.BitSwapsRequiered(29, 15)).Eql(2)
		})

		g.It("Should return 0 if both numbers are the same", func() {
			g.Assert(bits.BitSwapsRequiered(15, 15)).Eql(0)
		})

	})
}
