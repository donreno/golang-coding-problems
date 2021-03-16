package googleinterview_test

import (
	gi "golang-coding-problems/internal/googleinterview"

	"testing"

	"github.com/franela/goblin"
)

func TestStringSegmentation(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("String Segmentation", func() {
		dict := make(map[string]bool)
		dict["hello"] = true
		dict["on"] = true
		dict["now"] = true

		g.It("Should return true if string can be segmented in dictionary words", func() {
			g.Assert(gi.IsStringSegmentation("hellonow", dict)).IsTrue()
		})

		g.It("Should return true if string can not be segmented in dictionary words", func() {
			g.Assert(gi.IsStringSegmentation("helloqon", dict)).IsFalse()
		})
	})
}
