package crackingthecodinginterview_test

import (
	ctci "golang-coding-problems/internal/crackingthecodinginterview"

	"testing"

	"github.com/franela/goblin"
)

func TestMergeSort(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Merge Sort", func() {
		var arr []int

		g.It("Should sort elements properly", func() {
			ctci.Mergesort(arr)
			g.Assert(arr).Eql([]int{2, 11, 15, 28, 43, 66, 91})
		})

		g.BeforeEach(func() {
			arr = []int{91, 43, 11, 2, 28, 66, 15}
		})
	})
}
