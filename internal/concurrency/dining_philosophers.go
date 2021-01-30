package concurrency

import (
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

const bitesToFinish = 5

// Philosopher represents a philosopher for Dinning Philosophers
type Philosopher int

// Think Philosopher is thinking
func (p Philosopher) think() {
	fmt.Printf("Philosopher %d is thinking...", int(p))
	time.Sleep(time.Millisecond * 100)
}

// Eat Philosopher is eating
func (p Philosopher) eat() {
	fmt.Printf("Philosopher %d is thinking...", int(p))
	time.Sleep(time.Millisecond * 100)
}

// Philosopher tries to get chopsticks
func (p Philosopher) getChopsticks(left, right *semaphore.Weighted) bool {
	fmt.Printf("Philosopher %d is going to grab chopsticks...", int(p))

	fmt.Printf("Philosopher %d is going to grab left chopstick...", int(p))
	if !left.TryAcquire(1) {
		fmt.Printf("Philosopher %d failed to grab left chopstick!", int(p))
		return false
	}

	if !right.TryAcquire(1) {
		fmt.Printf("Philosopher %d failed to grab right chopstick!", int(p))
		left.Release(1)
		return false
	}

	return true
}

func (p Philosopher) putChopsticksDown(left, right *semaphore.Weighted) {
	left.Release(1)
	right.Release(1)
}

// Dine Philosopher is hungry and needs to eat and think
func (p Philosopher) Dine(left, right *semaphore.Weighted) {
	for bites := 0; bites <= bitesToFinish; {
		p.think()

		if p.getChopsticks(left, right) {
			p.eat()
			p.putChopsticksDown(left, right)
			bites++
		}
	}
}
