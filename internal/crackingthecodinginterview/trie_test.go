package crackingthecodinginterview_test

import (
	ctci "golang-coding-problems/internal/crackingthecodinginterview"
	"testing"

	"github.com/franela/goblin"
)

func TestTrie(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Trie Data struct", func() {
		var trie *ctci.Trie
		var words []string

		g.It("Should find all added words", func() {
			for _, word := range words {
				g.Assert(trie.HasWord(word)).IsTrue()
			}
		})

		g.It("Should not find a word that's not on trie", func() {
			g.Assert(trie.HasWord("nope")).IsFalse()
		})

		g.BeforeEach(func() {
			trie = ctci.MakeTrie()
			words = []string{
				"hello",
				"help",
				"land",
				"love",
			}

			trie.AddWords(words)
		})

	})
}
