package recursionndynamic_test

import (
	r "golang-coding-problems/internal/recursionndynamic"
	"testing"

	goblin "github.com/franela/goblin"
)

func TestFibonacci(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Fibonacci implementations", func() {

		g.It("Should return nth terms with Memo Fibo", func() {
			g.Assert(r.MemoFibonacci(0)).Eql(0)
			g.Assert(r.MemoFibonacci(1)).Eql(1)
			g.Assert(r.MemoFibonacci(2)).Eql(1)
			g.Assert(r.MemoFibonacci(3)).Eql(2)
			g.Assert(r.MemoFibonacci(4)).Eql(3)
			g.Assert(r.MemoFibonacci(5)).Eql(5)
			g.Assert(r.MemoFibonacci(6)).Eql(8)
		})

		g.It("Should return nth terms with Bottom Up Fibo", func() {
			g.Assert(r.BottomupDPFibonacci(0)).Eql(0)
			g.Assert(r.BottomupDPFibonacci(1)).Eql(1)
			g.Assert(r.BottomupDPFibonacci(2)).Eql(1)
			g.Assert(r.BottomupDPFibonacci(3)).Eql(2)
			g.Assert(r.BottomupDPFibonacci(4)).Eql(3)
			g.Assert(r.BottomupDPFibonacci(5)).Eql(5)
			g.Assert(r.BottomupDPFibonacci(6)).Eql(8)
		})

	})
}

func BenchmarkMemoFibo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.MemoFibonacci(n)
	}
}

func BenchmarkBottomUpFibo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r.BottomupDPFibonacci(n)
	}
}
