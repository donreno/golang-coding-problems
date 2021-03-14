package googleinterview_test

import (
	gi "golang-coding-problems/internal/googleinterview"
	"testing"

	"github.com/franela/goblin"
)

func TestSumOfTwo(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Check if Array containts two values that sum a target value", func() {
		var sumOfTwo gi.SumOfTwo

		g.Describe("Considering array unsorted", func() {
			sumOfTwo = gi.SumOfTwoUnsorted

			g.It("Should return false if array doesnt contain two values that sum target", func() {
				g.Assert(sumOfTwo([]int{2, 90, 11}, 3)).IsFalse()
			})

			g.It("Should return true if array contains two values that sum target", func() {
				g.Assert(sumOfTwo([]int{1, 2, 3}, 4)).IsTrue()
			})
		})

		g.Describe("Considering array sorted", func() {
			sumOfTwo = gi.SumOfTwoSorted

			g.It("Should return false if array doesnt contain two values that sum target", func() {
				g.Assert(sumOfTwo([]int{2, 11, 90}, 3)).IsFalse()
			})

			g.It("Should return true if array contains two values that sum target", func() {
				g.Assert(sumOfTwo([]int{1, 2, 3}, 4)).IsTrue()
			})
		})
	})
}
