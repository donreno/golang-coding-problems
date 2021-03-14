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
	keys := []string{}

	for _, s := range arr {
		chars := characters([]rune(s))
		sort.Sort(chars)
		key := string(chars)

		if _, found := anagramsMap[key]; !found {
			anagramsMap[key] = new(structs.LinkedList)
			keys = append(keys, key)
		}

		anagramsMap[key].Add(s)
	}

	sort.Strings(keys)

	i := 0
	for _, key := range keys {
		for _, elem := range anagramsMap[key].ToSlice() {
			s := elem.(string)

			arr[i] = s
			i++
		}
	}

}
