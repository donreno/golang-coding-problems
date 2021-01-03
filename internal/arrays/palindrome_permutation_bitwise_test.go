package arrays_test

import (
	a "golang-coding-problems/internal/arrays"
	"testing"

	goblin "github.com/franela/goblin"
)

/*
Given a string, write a function to check if it is a permutation of a palin-
drome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation
is a rea rrangement of letters. The palindrome does not need to be limited to just dictionary words.
*/

func TestIsPalindromePermutationBitwise(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Palindrome Permutations (Bitwise)", func() {
		g.It("Should return false if string is empty", func() {
			g.Assert(a.IsPalindromePermutationBitwise("")).IsFalse()
		})

		g.It("Should return false if string is not a Palindrome permutation", func() {
			g.Assert(a.IsPalindromePermutationBitwise("blaob")).IsFalse()
		})

		g.It("Should return true if string is a Palindrome permutation", func() {
			g.Assert(a.IsPalindromePermutationBitwise("tact coa")).IsTrue()
			g.Assert(a.IsPalindromePermutationBitwise("Tact Coa")).IsTrue()
			g.Assert(a.IsPalindromePermutationBitwise("cargofargoc")).IsTrue()
		})
	})
}

func BenchmarkPalindromePermutationsBitwise(b *testing.B) {

	for n := 0; n < b.N; n++ {
		a.IsPalindromePermutationBitwise("cargofargoc")
	}
}
