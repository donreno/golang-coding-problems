package golang_coding_problems_test

import (
	c "golang-coding-problems"
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
			g.Assert(c.CheckBalancedString("[{]")).IsFalse()
		})

		g.It("Should return false if string has chars in between is not balanced", func() {
			g.Assert(c.CheckBalancedString("AA]")).IsFalse()
		})

		g.It("Should return true if string is balanced", func() {
			g.Assert(c.CheckBalancedString("[{}]")).IsTrue()
		})

		g.It("Should return true if string is balanced and there is text in between balanced", func() {
			g.Assert(c.CheckBalancedString("[A{HI}]ASD")).IsTrue()
		})

		g.It("Should return true if string is empty", func() {
			g.Assert(c.CheckBalancedString("")).IsTrue()
		})
	})
}

func BenchmarkBalancedString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		c.CheckBalancedString("[{}]")
	}
}
