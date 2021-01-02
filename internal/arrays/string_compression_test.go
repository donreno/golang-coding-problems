package arrays_test

import (
	"testing"

	goblin "github.com/franela/goblin"

	a "golang-coding-problems/internal/arrays"
)

func TestStringsCompression(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Strings compression", func() {
		g.It("Should return same string if len is lower or equal to 2", func() {
			g.Assert(a.CompressString("aa")).Eql("aa")
			g.Assert(a.CompressString("")).Eql("")
		})

		g.It("Should return same string if len is lower than compressed len", func() {
			g.Assert(a.CompressString("abc")).Eql("abc")
		})

		g.It("Should return compressed string if len is lower than original", func() {
			g.Assert(a.CompressString("aaaabbcaaa")).Eql("a4b2c1a3")
		})
	})
}
