package structs_test

import (
	s "golang-coding-problems/internal/structs"
	"testing"

	goblin "github.com/franela/goblin"
)

type testStruct struct {
	name string
}

func TestHashtable(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("Hashtable", func() {
		var hashtable s.Hashtable

		g.BeforeEach(func() {
			hashtable = s.MakeHashtable()
		})

		g.It("Should return nil if element to get is not in table", func() {
			hashtable.Put("hello", "hello")

			g.Assert(hashtable.Get("zzzzz")).IsNil()
			g.Assert(hashtable.Get("hello1")).IsNil()
		})

		g.It("Should put elements properly", func() {
			dog := &testStruct{"dog"}
			hashtable.Put("dog", dog)

			g.Assert(hashtable.Get("dog")).Eql(dog)
		})

		g.It("Should be able to put multiple elements then retrieve all of them", func() {
			dog := &testStruct{"dog"}
			cat := &testStruct{"cat"}
			beer := &testStruct{"beer"}
			zzzzz := &testStruct{"zzzzz"}

			hashtable.Put("dog", dog)
			hashtable.Put("cat", cat)
			hashtable.Put("beer", beer)
			hashtable.Put("zzzzzzzz", zzzzz)

			g.Assert(hashtable.Get("dog")).Eql(dog)
			g.Assert(hashtable.Get("cat")).Eql(cat)
			g.Assert(hashtable.Get("beer")).Eql(beer)
			g.Assert(hashtable.Get("zzzzzzzz")).Eql(zzzzz)
		})

		g.It("Should replace previous element with new one if the key already exists", func() {
			pizza := &testStruct{"pizza"}
			pepperoniPizza := &testStruct{"pepperoni pizza"}

			hashtable.Put("pizza", pizza)
			hashtable.Put("pizza", pepperoniPizza)

			g.Assert(hashtable.Get("pizza")).Eql(pepperoniPizza)
		})
	})
}

func BenchmarkHashtable(b *testing.B) {
	dog := &testStruct{"dog"}
	cat := &testStruct{"cat"}
	beer := &testStruct{"beer"}
	pizza := &testStruct{"pizza"}

	h := s.MakeHashtable()
	h.Put("dog", dog)
	h.Put("cat", cat)
	h.Put("beer", beer)

	for n := 0; n < b.N; n++ {
		h.Put("pizza", pizza)

		h.Get("beer")
		h.Get("asd")
		h.Get("dog")
		h.Get("cat")
		h.Get("other")
		h.Get("pizza")
	}
}
