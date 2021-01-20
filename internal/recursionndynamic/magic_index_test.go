package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"

	"testing"

	goblin "github.com/franela/goblin"
)

func TestMagicIndex(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Magic Index", func() {
		g.It("Should identify that array is magic index", func() {

			arr := []int{-1, 0, 1, 3, 5, 7, 9}

			g.Assert(r.HasMagicindex(arr)).IsTrue()
		})

		g.It("Should identify that array is not magic index", func() {

			arr := []int{-1, 0, 1, 2, 5, 7, 9}

			g.Assert(r.HasMagicindex(arr)).IsFalse()
		})

		g.It("Should return false if array is empty", func() {

			arr := []int{}

			g.Assert(r.HasMagicindex(arr)).IsFalse()
		})

		g.It("Should identify that array is magic index if all elements are the same", func() {

			arr := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}

			g.Assert(r.HasMagicindex(arr)).IsTrue()
		})
	})

}
