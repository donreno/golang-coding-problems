package arrays_test

import (
	"testing"

	a "golang-coding-problems/internal/arrays"

	goblin "github.com/franela/goblin"
)

/*
There are three types of edits that can be performed on strings: insert a character,
remove a character, or replace a character. Given two strings, write a function to check if they are
one edit (or zero edits) away.
*/

func TestOneAway(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("One Away", func() {
		g.It("Should return false if both strings are not one change away", func() {
			g.Assert(a.IsOneAway("abc", "abcgh")).IsFalse()
			g.Assert(a.IsOneAway("abc", "hij")).IsFalse()
			g.Assert(a.IsOneAway("pale", "bake")).IsFalse()
		})

		g.It("Should return true if both strings are zero or one change away", func() {
			g.Assert(a.IsOneAway("pale", "ple")).IsTrue()
			g.Assert(a.IsOneAway("pales", "pale")).IsTrue()
			g.Assert(a.IsOneAway("pale", "pale")).IsTrue()
		})
	})
}
