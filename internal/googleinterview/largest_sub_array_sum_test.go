package googleinterview_test

import (
	gi "golang-coding-problems/internal/googleinterview"

	"testing"

	"github.com/franela/goblin"
)

func TestLargestSubArraySum(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Largest sub array sum", func() {
		arr := []int{-4, 2, -5, 1, 2, 3, 6, -5, 1}

		g.It("Should find right largest sub array sum", func() {
			g.Assert(gi.FindLargestSubArraySum(arr)).Equal(12)
		})
	})
}
