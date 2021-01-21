package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"

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
			g.Assert(perms[0]).Eql("AB")
			g.Assert(perms[1]).Eql("BA")
		})

		g.It("Should find all permutations for three elements case", func() {
			perms := r.PermutationsWithDups(threeElementsCase)
			g.Assert(perms[0]).Eql("AAC")
			g.Assert(perms[1]).Eql("ACA")
			g.Assert(perms[2]).Eql("CAA")
		})
	})

}
