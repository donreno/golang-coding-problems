package stacks_test

import (
	s "golang-coding-problems/internal/stacks"
	"testing"

	goblin "github.com/franela/goblin"
)

/*
An animal shelter, which holds only dogs and cats, operates on a strictly"first in, first
out" basis. People must adopt either the "oldest" (based on arrival time) of all animals at the shelter,
or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of
that type). They cannot select which specific animal they would like. Create the data structures to
maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog,
and dequeueCat. You may use the built-in Linked List data structure.
*/
func TestAnimalShelter(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Animal Shelter", func() {
		var animalShelter s.AnimalShelter

		g.BeforeEach(func() {
			animalShelter = s.MakeAnimalShelter()
		})

		g.It("Should behave as described on top", func() {
			animalShelter.Enqueue(&s.Pet{s.Cat})
			animalShelter.Enqueue(&s.Pet{s.Dog})
			animalShelter.Enqueue(&s.Pet{s.Dog})
			animalShelter.Enqueue(&s.Pet{s.Cat})
			animalShelter.Enqueue(&s.Pet{s.Dog})

			g.Assert(animalShelter.DequeueDog().Kind).Eql(s.Dog)
			g.Assert(animalShelter.DequeueAny().Kind).Eql(s.Cat)
			g.Assert(animalShelter.DequeueCat().Kind).Eql(s.Cat)
			g.Assert(animalShelter.DequeueAny().Kind).Eql(s.Dog)
			g.Assert(animalShelter.DequeueAny().Kind).Eql(s.Dog)
			g.Assert(animalShelter.Empty()).IsTrue()
		})
	})

}
