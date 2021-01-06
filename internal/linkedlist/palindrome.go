package linkedlist

import (
	"bytes"
	s "golang-coding-problems/internal/structs"
)

// IsPalindrome checks if linkedlist is palindrome
func IsPalindrome(ll *s.LinkedList) bool {
	if ll.Empty() {
		return false
	}

	// b - u - s - a - s - u - b

	word := new(bytes.Buffer)

	for current := ll.Head; current != nil; current = current.Next {
		word.WriteRune(current.Val.(rune))
	}

	wordRunes := []rune(word.String())
	currentNode := ll.Head

	for i := len(wordRunes) - 1; i >= len(wordRunes)/2; i, currentNode = i-1, currentNode.Next {
		if wordRunes[i] != currentNode.Val {
			return false
		}
	}

	return true
}
