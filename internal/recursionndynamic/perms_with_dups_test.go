package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"

	"github.com/stretchr/stew/slice"

	"testing"

	goblin "github.com/franela/goblin"
)

func TestPermutationsWithDups(t *testing.T) {
	g := goblin.Goblin(t)

	baseCase := "AB"
	threeElementsCase := "AAC"

	g.Describe("Permutations With Dups", func() {

		g.It("Should find all permutations for base case", func() {
			perms := r.PermutationsWithDups(baseCase)
			g.Assert(slice.Contains(perms, "AB")).IsTrue()
			g.Assert(slice.Contains(perms, "BA")).IsTrue()
			g.Assert(len(perms)).Eql(2)
		})

		g.It("Should find all permutations for three elements case", func() {
			perms := r.PermutationsWithDups(threeElementsCase)
			g.Assert(slice.Contains(perms, "AAC")).IsTrue()
			g.Assert(slice.Contains(perms, "ACA")).IsTrue()
			g.Assert(slice.Contains(perms, "CAA")).IsTrue()
			g.Assert(len(perms)).Eql(3)
		})
	})

}
