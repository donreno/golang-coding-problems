package sorting_test

import (
	s "golang-coding-problems/internal/sorting"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestGroupAnagrams(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Group Anagrams", func() {

		var arr []string

		g.BeforeEach(func() {
			arr = []string{"acre", "batman", "crea", "race", "tabman", "beer"}
		})

		g.It("Should put anagrams next to each other", func() {
			s.GroupAnagrams(arr)
			g.Assert(arr[0]).Eql("acre")
			g.Assert(arr[1]).Eql("crea")
			g.Assert(arr[2]).Eql("race")
			g.Assert(arr[3]).Eql("beer")
			g.Assert(arr[4]).Eql("batman")
			g.Assert(arr[5]).Eql("tabman")
		})
	})
}
