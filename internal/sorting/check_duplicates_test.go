package sorting_test

import (
	s "golang-coding-problems/internal/sorting"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestCheckDuplicates(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Check duplicates", func() {

		var arr []int

		g.BeforeEach(func() {
			arr = []int{1, 2, 3, 4, 2, 11, 27, 1, 65, 65, 3}
		})

		g.It("Should return expected map of duplicates", func() {
			dups := s.CheckDuplicates(arr)

			g.Assert(dups[1]).IsTrue()
			g.Assert(dups[2]).IsTrue()
			g.Assert(dups[3]).IsTrue()
			g.Assert(dups[65]).IsTrue()
			g.Assert(dups[27]).IsFalse()
		})
	})
}
