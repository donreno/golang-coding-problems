package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"

	"testing"

	goblin "github.com/franela/goblin"
)

func TestPermutationsWithoutDups(t *testing.T) {
	g := goblin.Goblin(t)

	baseCase := "AB"
	threeElementsCase := "ABC"

	g.Describe("Permutations Without Dups", func() {

		g.It("Should find all permutations for base case", func() {
			perms := r.PermutationsWithoutDups(baseCase)
			g.Assert(perms[0]).Eql("AB")
			g.Assert(perms[1]).Eql("BA")
		})

		g.It("Should find all permutations for base case", func() {
			perms := r.PermutationsWithoutDups(threeElementsCase)
			g.Assert(perms[0]).Eql("ABC")
			g.Assert(perms[1]).Eql("BAC")
			g.Assert(perms[2]).Eql("BCA")
			g.Assert(perms[3]).Eql("ACB")
			g.Assert(perms[4]).Eql("CAB")
			g.Assert(perms[5]).Eql("CBA")
		})
	})

}
