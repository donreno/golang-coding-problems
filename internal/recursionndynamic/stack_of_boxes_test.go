package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestStackOfBoxes(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Stack of Boxes", func() {
		var boxes []*r.Box

		g.It("Should return correct stack of boxes with greatest height", func() {
			g.Assert(r.StackOfBoxes(boxes)).Eql(14)
		})

		g.It("Should return height 0 if boxes list is empty", func() {
			boxes = []*r.Box{}
			g.Assert(r.StackOfBoxes(boxes)).Eql(0)
		})

		g.BeforeEach(func() {
			boxes = []*r.Box{
				&r.Box{
					Height: 3,
					Width:  3,
					Depth:  2,
				},
				&r.Box{
					Height: 7,
					Width:  12,
					Depth:  1,
				},
				&r.Box{
					Height: 5,
					Width:  5,
					Depth:  5,
				},
				&r.Box{
					Height: 10,
					Width:  8,
					Depth:  1,
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
