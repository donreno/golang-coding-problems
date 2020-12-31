package golang_coding_problems_test

import (
	"testing"

	goblin "github.com/franela/goblin"

	c "golang-coding-problems"
)

func TestCheckPermutation(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Check if B is permutation of A", func() {
		g.It("Should return false if length of both strings is not equal", func() {
			g.Assert(c.IsBpermutationOfA("a", "ab")).IsFalse()
		})

		g.It("Should return false if b is not permutation of a", func() {
			g.Assert(c.IsBpermutationOfA("abc", "bcd")).IsFalse()
		})

		g.It("Should return true if b is permutation of a", func() {
			g.Assert(c.IsBpermutationOfA("ABCD1", "B1DCA")).IsTrue()
		})
	})
}
