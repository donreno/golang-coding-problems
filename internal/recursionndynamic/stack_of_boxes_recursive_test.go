package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestStackOfBoxesRecursive(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Stack of Boxes Recursive", func() {
		var boxes []*r.Box

		g.It("Should return correct stack of boxes with greatest height", func() {
			g.Assert(r.StackOfBoxesRecursive(boxes)).Eql(9)
		})

		g.BeforeEach(func() {
			boxes = []*r.Box{
				&r.Box{
					Height: 3,
					Width:  3,
					Depth:  2,
				},
				&r.Box{
					Height: 5,
					Width:  5,
					Depth:  5,
				},
				&r.Box{
					Height: 4,
					Width:  5,
					Depth:  5,
				},
				&r.Box{
					Height: 6,
					Width:  4,
					Depth:  3,
				},
			}
		})
	})
}
