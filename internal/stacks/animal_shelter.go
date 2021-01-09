package stacks

import (
	s "golang-coding-problems/internal/structs"
)

const (
	// Dog kind
	Dog = "dog"
	// Cat kind
	Cat = "cat"
)

// AnimalShelter interface to represent an animal shelter
type AnimalShelter interface {
	Enqueue(*Pet)
	DequeueAny() *Pet
	DequeueDog() *Pet
	DequeueCat() *Pet
	Empty() bool
}

// Pet thats gonna arrive or be adopted at shelter
type Pet struct {
	Kind string
}

type animalShelter struct {
	pets s.Queue
}

// MakeAnimalShelter creates animal shelter
func MakeAnimalShelter() AnimalShelter {
	return &animalShelter{
		pets: s.MakeQueue(),
	}
}

func (a *animalShelter) Empty() bool {
	return a.pets.Empty()
}

func (a *animalShelter) Enqueue(p *Pet) {
	a.pets.Enqueue(p)
}

func (a *animalShelter) DequeueAny() *Pet {
	pet := a.pets.Dequeue()

	if pet == nil {
		return nil
	}

	return pet.(*Pet)
}

func (a *animalShelter) DequeueDog() *Pet {
	return a.dequeueByKind(Dog)
}

func (a *animalShelter) DequeueCat() *Pet {
	return a.dequeueByKind(Cat)
}

// Will try to find oldest animal by kind, and dequeue it
func (a *animalShelter) dequeueByKind(kind string) *Pet {
	tempQueue := s.MakeQueue()

	var pet *Pet

	for !a.pets.Empty() {
		currentPet := a.pets.Dequeue().(*Pet)
		if currentPet.Kind == kind && pet == nil {
			pet = currentPet
		} else {
			tempQueue.Enqueue(currentPet)
		}
	}

	a.pets = tempQueue

	return pet
}
