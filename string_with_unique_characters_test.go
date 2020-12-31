package golang_coding_problems_test

import (
	"testing"

	goblin "github.com/franela/goblin"

	c "golang-coding-problems"
)

// Implement an algorithm to determine if a string has all unique characters
func TestStringWithUniqueCharacters(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("String With Unique Characters (Solved with structs [map])", func() {

		g.It("Should return false if string is empty", func() {
			g.Assert(c.HasUniqueChars("")).IsFalse()
		})

		g.It("Should return false if string is not unique", func() {
			g.Assert(c.HasUniqueChars("abcdefgaf1")).IsFalse()
		})

		g.It("Should return true if string is unique", func() {
			g.Assert(c.HasUniqueChars("abcdefghijklmnopqrstuvwxyz12345678")).IsTrue()
		})

		g.It("Should return true if string has just one char", func() {
			g.Assert(c.HasUniqueChars("1")).IsTrue()
		})
	})
}
