package golang_coding_problems

// CountPermutationsOfAonB counts permutations of A on B
func CountPermutationsOfAonB(a, b string) int {
	aLen := len(a)
	bLen := len(b)

	if bLen < 1 || aLen < 1 {
		return -1
	}

	if aLen > bLen {
		return -1
	}

	lastIndex := bLen - aLen
	count := 0
	aFreq := makeFreqMatrix(a)
	bRunes := []rune(b)

	for i := 0; i <= lastIndex; i++ {
		currentRuneSet := bRunes[i : i+aLen]
		currentFreq := makeFreqMatrix(string(currentRuneSet))

		if equalFreqMatrix(aFreq, currentFreq) {
			count++
		}
	}

	return count
}

func makeFreqMatrix(s string) map[rune]int {
	freq := make(map[rune]int)

	for _, c := range s {
		freq[c] = freq[c] + 1
	}

	return freq
}

func equalFreqMatrix(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}

	for c, count := range a {
		if b[c] != count {
			return false
		}
	}

	return true
}
