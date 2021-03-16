package googleinterview_test

import (
	gi "golang-coding-problems/internal/googleinterview"

	"testing"

	"github.com/franela/goblin"
)

func TestFindPalindromeSubstring(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Find Palindrome Substring", func() {
		word := "marcopolo"

		g.It("Should find right palindrome substrings", func() {
			palindromes := gi.FindAllPalindromeWords(word)
			g.Assert(palindromes[0]).Eql("olo")
			g.Assert(palindromes[1]).Eql("opo")
		})
	})
}
