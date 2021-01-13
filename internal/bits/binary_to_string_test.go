package bits_test

import (
	"golang-coding-problems/internal/bits"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestBinaryToString(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Binary to String", func() {

		g.It("Should convert number to binary correctly", func() {
			g.Assert(bits.BinaryToString(0.75)).Eql("0.11")
			g.Assert(bits.BinaryToString(0.25)).Eql("0.01")
			g.Assert(bits.BinaryToString(0.125)).Eql("0.001")
		})

		g.It("Should fail if number is greater than 32 bits", func() {
			g.Assert(bits.BinaryToString(1.1)).Eql("ERROR")
			g.Assert(bits.BinaryToString(0.1)).Eql("ERROR") // 0.0001100110011001100110011001100110011001100110011001101
			g.Assert(bits.BinaryToString(0.999999999999999999)).Eql("ERROR")
		})

	})
}
