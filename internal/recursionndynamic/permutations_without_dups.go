package recursionndynamic

import (
	"strings"
)

// PermutationsWithoutDups returns all permutations
func PermutationsWithoutDups(s string) []string {
	if s == "" {
		return []string{""}
	}

	permutations := []string{}

	sAsRunes := []rune(s)
	firstRune := sAsRunes[0]

	words := PermutationsWithoutDups(string(sAsRunes[1:]))

	for i := 0; i < len(words); i++ {
		runeWord := []rune(words[i])
		for j := 0; j <= len(runeWord); j++ {
			strBuild := new(strings.Builder)

			strBuild.WriteString(string(runeWord[:j]))
			strBuild.WriteRune(firstRune)
			strBuild.WriteString(string(runeWord[j:]))
			permutations = append(permutations, strBuild.String())
		}
	}

	return permutations
}
