package crackingthecodinginterview_test

import (
	ctci "golang-coding-problems/internal/crackingthecodinginterview"

	"testing"

	"github.com/franela/goblin"
)

func TestIcecreamParlor(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Icecream Parlor", func() {
		var prices []int
		var money int

		g.It("Should find expected icecream indices (without binary search", func() {
			g.Assert(ctci.GetIcecreamChoicesBS(prices, money)).Eql([]int{3, 6})
		})

		g.It("Should find expected icecream indices (with binary search)", func() {
			g.Assert(ctci.GetIcecreamChoicesBS(prices, money)).Eql([]int{3, 6})
		})

		g.BeforeEach(func() {
			prices = []int{2, 7, 13, 5, 4, 13, 5}
			money = 10
		})
	})
}
