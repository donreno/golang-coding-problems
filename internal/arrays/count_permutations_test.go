package arrays_test

import (
	"testing"

	goblin "github.com/franela/goblin"

	a "golang-coding-problems/internal/arrays"
)

// e.g: if A is "abc" and B is "acbdbacaf1abc", then there are 3 permutations of A on B
func TestCountPermutations(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Count Permutations Test", func() {

		g.It("Should return -1 if A is longer than B", func() {
			perms := a.CountPermutationsOfAonB("abc", "a")

			g.Assert(perms).Eql(-1)
		})

		g.It("Should return -1 if A is empty", func() {
			perms := a.CountPermutationsOfAonB("", "a")

			g.Assert(perms).Eql(-1)
		})

		g.It("Should return -1 if B is empty", func() {
			perms := a.CountPermutationsOfAonB("a", "")

			g.Assert(perms).Eql(-1)
		})

		g.It("Should find the right number of permutations on given strings", func() {
			perms := a.CountPermutationsOfAonB("abc", "cbabadcbbababaabccbac")

			g.Assert(perms).Eql(4)
		})
	})
}

func BenchmarkCountPermutations(b *testing.B) {

	for n := 0; n < b.N; n++ {
		a.CountPermutationsOfAonB("abcd", "cbcbabadcbbababaabccbacabacbabadcbbababaabccbacdccbabadcbbababaabccbacbbababaabccbcbabadcbbababaabccbacac")
	}
}
