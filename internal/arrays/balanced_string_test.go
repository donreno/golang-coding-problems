package arrays_test

import (
	a "golang-coding-problems/internal/arrays"
	"testing"

	goblin "github.com/franela/goblin"
)

/*
* Balanced string problem:
* Checks if string's parentheses are balanced
* e.g: [{()}] is balanced, and {[(]} is not balanced
 */
func TestBalancedString(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Balanced string problem", func() {
		g.It("Should return false if string is not balanced", func() {
			g.Assert(a.CheckBalancedString("[{]")).IsFalse()
		})

		g.It("Should return false if string has chars in between is not balanced", func() {
			g.Assert(a.CheckBalancedString("AA]")).IsFalse()
		})

		g.It("Should return true if string is balanced", func() {
			g.Assert(a.CheckBalancedString("[{}]")).IsTrue()
		})

		g.It("Should return true if string is balanced and there is text in between balanced", func() {
			g.Assert(a.CheckBalancedString("[A{HI}]ASD")).IsTrue()
		})

		g.It("Should return true if string is empty", func() {
			g.Assert(a.CheckBalancedString("")).IsTrue()
		})
	})
}
