package sorting

import (
	"golang-coding-problems/internal/structs"
	"sort"
)

type characters []rune

func (c characters) Len() int {
	return len(c)
}

func (c characters) Less(i, j int) bool {
	return c[i] > c[j]
}

func (c characters) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// GroupAnagrams groups anagrams
func GroupAnagrams(arr []string) {
	anagramsMap := make(map[string]*structs.LinkedList)

	for _, s := range arr {
		chars := characters([]rune(s))
		sort.Sort(chars)
		key := string(chars)

		if _, found := anagramsMap[key]; !found {
			anagramsMap[key] = new(structs.LinkedList)
		}

		anagramsMap[key].Add(s)
	}

	i := 0
	for _, l := range anagramsMap {
		for _, elem := range l.ToSlice() {
			s := elem.(string)

			arr[i] = s
			i++
		}
	}

}
