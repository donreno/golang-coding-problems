package stacks_test

import (
	s "golang-coding-problems/internal/stacks"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestStacks(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Three in one", func() {
		var threeInOne *s.ThreeStacksInOneArray

		g.BeforeEach(func() {
			threeInOne = s.MakeThreeInOneStack(3)
		})

		g.It("Should handle 3 stacks in one array", func() {
			threeInOne.Push(15, 0)
			threeInOne.Push(10, 1)
			threeInOne.Push(11, 1)
			threeInOne.Push(20, 2)

			firstStackPeek, _ := threeInOne.Peek(0)
			secondStackPop, _ := threeInOne.Pop(1)
			secondStackPeek, _ := threeInOne.Peek(1)
			thirdStackPop, _ := threeInOne.Pop(2)

			g.Assert(firstStackPeek).Eql(15)
			g.Assert(secondStackPop).Eql(11)
			g.Assert(secondStackPeek).Eql(10)
			g.Assert(thirdStackPop).Eql(20)
		})

		g.It("Should fail if I peek on empty stack", func() {

			_, err := threeInOne.Peek(0)

			g.Assert(err).IsNotNil()
		})

		g.It("Should fail if I pop on empty stack", func() {
			_, err := threeInOne.Pop(0)

			g.Assert(err).IsNotNil()
		})

		g.It("Should fail if I push on invalid stack number", func() {
			err := threeInOne.Push(5, 5)

			g.Assert(err).IsNotNil()
		})

		g.It("Should fail if I overflow stack pushing more than it's fixed size", func() {
			threeInOne.Push(1, 0)
			threeInOne.Push(2, 0)
			threeInOne.Push(3, 0)

			err := threeInOne.Push(4, 0)

			g.Assert(err).IsNotNil()
		})

		g.It("Should find Min element on stack", func() {
			threeInOne.Push(23, 0)
			threeInOne.Push(2, 0)
			threeInOne.Push(999, 0)

			min, _ := threeInOne.Min(0)

			g.Assert(min).Eql(2)
		})
	})
}
