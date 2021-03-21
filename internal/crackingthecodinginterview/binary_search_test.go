package crackingthecodinginterview_test

import (
	ctci "golang-coding-problems/internal/crackingthecodinginterview"

	"testing"

	"github.com/franela/goblin"
)

func TestBinarySearch(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Binary Search", func() {
		var arr []int

		g.It("Recursive: Should return true if element is on array", func() {
			g.Assert(ctci.BinarySearchRecursive(arr, 121)).IsTrue()
		})

		g.It("Recursive: Should return false if element is on array", func() {
			g.Assert(ctci.BinarySearchRecursive(arr, 99)).IsFalse()
		})

		g.It("Non Recursive: Should return true if element is on array", func() {
			g.Assert(ctci.BinarySearchNonRecursive(arr, 121)).IsTrue()
		})

		g.It("Non Recursive: Should return false if element is on array", func() {
			g.Assert(ctci.BinarySearchNonRecursive(arr, 99)).IsFalse()
		})

		g.Before(func() {
			arr = []int{0, 3, 8, 11, 12, 15, 18, 20, 21, 22, 27, 45, 121, 130, 131, 150}
		})
	})
}
