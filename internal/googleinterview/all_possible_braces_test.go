package googleinterview_test

import (
	gi "golang-coding-problems/internal/googleinterview"

	"testing"

	"github.com/franela/goblin"
)

func TestAllPossibleBraces(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("All possible braces", func() {
		expectedBraces := []string{
			"{{{}}}",
			"{{}{}}",
			"{{}}{}",
			"{}{{}}",
			"{}{}{}",
		}

		g.It("Should generate all expected braces", func() {
			braces := gi.AllPossibleBraces(3)

			for i := range expectedBraces {
				g.Assert(braces[i]).Eql(expectedBraces[i])
			}
		})
	})
}
