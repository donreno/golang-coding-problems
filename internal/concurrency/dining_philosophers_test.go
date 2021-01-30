package concurrency_test

import (
	c "golang-coding-problems/internal/concurrency"
	"testing"
	"time"

	"golang.org/x/sync/semaphore"

	goblin "github.com/franela/goblin"
)

func TestDiningPhilosophers(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Dining Philosophers", func() {
		var philosophers []c.Philosopher
		var left *semaphore.Weighted
		var right *semaphore.Weighted

		g.It("Will let philosophers dine", func() {
			g.Timeout(time.Second * 10)

			for _, p := range philosophers {
				go p.Dine(left, right)
			}

			time.Sleep(time.Second * 5)
			g.Assert(left.TryAcquire(1)).Eql(true)
			g.Assert(right.TryAcquire(1)).Eql(true)
		})

		g.AfterEach(func() {
			left.Release(1)
			right.Release(1)
		})

		g.Before(func() {
			philosophers = []c.Philosopher{c.Philosopher(1), c.Philosopher(2), c.Philosopher(3)}
			left = semaphore.NewWeighted(1)
			right = semaphore.NewWeighted(1)
		})
	})
}
