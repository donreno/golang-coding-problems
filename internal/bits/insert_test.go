package bits_test

import (
	"golang-coding-problems/internal/bits"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestInsertBits(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Bit inserts", func() {

		g.It("Should insert properly M into N at right position", func() {
			n := 1024        // 10000000000
			m := 19          // 10011
			expected := 1100 // 10001001100

			g.Assert(bits.InsertMIntoNAtItoJ(m, n, 2, 6)).Eql(expected)
		})

	})
}
