package linkedlist_test

import (
	l "golang-coding-problems/internal/linkedlist"
	s "golang-coding-problems/internal/structs"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestPalindrome(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Check Linkedlist is Palindrome", func() {
		g.It("Should return true if linkedlist is palindrome", func() {

			ll := new(s.LinkedList).
				Add('s').Add('u').Add('b').
				Add('a').
				Add('b').Add('u').Add('s')

			g.Assert(l.IsPalindrome(ll)).IsTrue()
		})

		g.It("Should return false if linkedlist is not palindrome", func() {
			ll := new(s.LinkedList).
				Add('s').Add('o').Add('b').Add('a').
				Add('b').Add('u').Add('s')

			g.Assert(l.IsPalindrome(ll)).IsFalse()
		})
	})
}
