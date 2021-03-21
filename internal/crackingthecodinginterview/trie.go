package crackingthecodinginterview

// Trie data struct
type Trie struct {
	Child  map[rune]*Trie
	IsWord bool
}

// MakeTrie makes a new trie
func MakeTrie() *Trie {
	return &Trie{
		Child: make(map[rune]*Trie),
	}
}

// AddWords add multiple words to trie
func (t *Trie) AddWords(words []string) {

	for _, word := range words {
		t.AddWord(word)
	}
}

// AddWord adds single word to trie
func (t *Trie) AddWord(word string) {
	current := t

	for _, letter := range []rune(word) {
		if currentLetterTrie, found := current.Child[letter]; found {
			current = currentLetterTrie
		} else {
			current.Child[letter] = MakeTrie()
			current = current.Child[letter]
		}
	}

	current.IsWord = true
}

// HasWord checks if trie has given word
func (t *Trie) HasWord(word string) bool {
	current := t
	for _, letter := range []rune(word) {
		if _, found := current.Child[letter]; !found {
			return false
		}

		current = current.Child[letter]
	}

	return current.IsWord
}
