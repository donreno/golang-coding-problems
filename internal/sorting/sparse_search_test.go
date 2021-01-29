package sorting_test

import (
	s "golang-coding-problems/internal/sorting"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestSparseSearch(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Sparse search", func() {

		var strs []string

		g.BeforeEach(func() {
			strs = []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
		})

		g.It("Should find expected element index", func() {
			g.Assert(s.SparseSearch(strs, "ball")).Eql(4)
			g.Assert(s.SparseSearch(strs, "at")).Eql(0)
			g.Assert(s.SparseSearch(strs, "cat")).Eql(-1)
		})

	})
}
