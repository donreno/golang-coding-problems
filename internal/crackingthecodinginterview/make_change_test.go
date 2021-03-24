package crackingthecodinginterview_test

import (
	ctci "golang-coding-problems/internal/crackingthecodinginterview"

	"testing"

	"github.com/franela/goblin"
)

func TestMakeChange(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Make Change (Count ways to make change)", func() {
		var money int
		var coins []int

		g.It("Should return expected number of ways to make money", func() {
			ways := ctci.MakeChange(coins, money)

			g.Assert(ways).Eql(13)
		})

		g.BeforeEach(func() {
			money = 27
			coins = []int{25, 10, 5, 1}
		})
	})
}
