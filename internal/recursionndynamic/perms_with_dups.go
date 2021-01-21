package recursionndynamic

import (
	"strings"
)

// PermutationsWithDups permutations with duplicates
func PermutationsWithDups(s string) []string {
	return permsWithDups(buildFreqMatrixForPerms(s), "", len(s))
}

func permsWithDups(matrix map[rune]int, prefix string, remaining int) []string {
	if remaining == 0 {
		return []string{prefix}
	}

	perms := []string{}

	for c, count := range matrix {
		if count > 0 {
			strBuild := new(strings.Builder)
			strBuild.WriteString(prefix)
			strBuild.WriteRune(c)
			matrix[c]--
			perms = append(perms, permsWithDups(matrix, strBuild.String(), remaining-1)...)
			matrix[c]++
		}
	}

	return perms
}

func buildFreqMatrixForPerms(s string) map[rune]int {
	m := make(map[rune]int)

	for _, r := range s {
		m[r]++
	}

	return m
}
