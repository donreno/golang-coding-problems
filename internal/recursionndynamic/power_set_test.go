package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"

	"testing"

	goblin "github.com/franela/goblin"
)

func TestPowerset(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("Power Set", func() {

		g.It("Should create expected power sets from a set of length 2", func() {
			set := []int{1, 2}
			subsets := r.GetAllSubsets(set)

			g.Assert(subsets[0]).Eql([]int{})
			g.Assert(subsets[1]).Eql([]int{2})
			g.Assert(subsets[2]).Eql([]int{1})
			g.Assert(subsets[3]).Eql([]int{1, 2})
		})

		g.It("Should create expected power sets from a set of length 3", func() {
			set := []int{1, 2, 3}
			subsets := r.GetAllSubsets(set)

			g.Assert(subsets[0]).Eql([]int{})
			g.Assert(subsets[1]).Eql([]int{3})
			g.Assert(subsets[2]).Eql([]int{2})
			g.Assert(subsets[3]).Eql([]int{2, 3})
			g.Assert(subsets[4]).Eql([]int{1})
			g.Assert(subsets[5]).Eql([]int{1, 3})
			g.Assert(subsets[6]).Eql([]int{1, 2})
			g.Assert(subsets[7]).Eql([]int{1, 2, 3})

		})
	})

}
