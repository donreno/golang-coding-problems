package golang_coding_problems_test

import (
	"testing"

	goblin "github.com/franela/goblin"

	c "golang-coding-problems"
)

/*
Write a method to replace all spaces in a string with '%20': You may assume that the string
has sufficient space at the end to hold the additional characters, and that you are given the "true"
length of the string. (Note: If implementing in Java, please use a character array so that you can
perform this operation in place.)
*/

func TestUrlify(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Urlify", func() {
		// will be using runes slice to check modifications after
		g.It("Should return false if string is empty", func() {
			runes := []rune("")
			g.Assert(c.Urlify(runes, 0)).IsFalse()
			g.Assert(runes).Equal([]rune(""))
		})

		g.It("Should return false if string has only whitespaces", func() {
			runes := []rune("   ")
			g.Assert(c.Urlify(runes, 0)).IsFalse()
			g.Assert(runes).Equal([]rune("   "))
		})

		g.It("Should return false if string does not have enough right spaces to admit all replacements", func() {
			runes := []rune("Mr John Smith   ")
			g.Assert(c.Urlify(runes, 16)).IsFalse()
			g.Assert(runes).Equal([]rune("Mr John Smith   "))
		})

		g.It("Should return true if string has enough right spaces to admit all replacements", func() {
			runes := []rune("Mr John Smith    ")
			g.Assert(c.Urlify(runes, 17)).IsTrue()
			g.Assert(runes).Equal([]rune("Mr%20John%20Smith"))
		})

		g.It("Should return true if string has enough right spaces to admit all replacements", func() {
			runes := []rune(" Mr John Smith      ")
			g.Assert(c.Urlify(runes, 20)).IsTrue()
			g.Assert(runes).Equal([]rune("%20Mr%20John%20Smith"))
		})
	})
}
