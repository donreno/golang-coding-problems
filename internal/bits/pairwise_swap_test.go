package bits_test

import (
	"golang-coding-problems/internal/bits"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestPairwiseSwap(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Pairwise swap", func() {

		g.It("Should properly swap even for odds", func() {
			g.Assert(bits.PairwiseSwap(7)).Eql(int16(11))
			g.Assert(bits.PairwiseSwap(3)).Eql(int16(3))
		})
	})
}
