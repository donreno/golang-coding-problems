package googleinterview

// AllPossibleBraces generates all possible braces
func AllPossibleBraces(len int) []string {
	possibleBraces := []string{}
	allPossibleBraces(&possibleBraces, []rune(""), 0, 0, len)
	return possibleBraces
}

func allPossibleBraces(possibleBraces *[]string, currentBraces []rune, open, close, max int) {
	if len(currentBraces) == max*2 {
		*possibleBraces = append(*possibleBraces, string(currentBraces))
	}

	if open < max {
		allPossibleBraces(possibleBraces, append(currentBraces, '{'), open+1, close, max)
	}

	if close < open {
		allPossibleBraces(possibleBraces, append(currentBraces, '}'), open, close+1, max)
	}
}
