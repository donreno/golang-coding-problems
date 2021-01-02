package arrays

var singleSpaceRune = rune(' ')
var urlSpace = []rune("%20")

// Urlify replaces non-left or right whitespaces with '%20' on given url
func Urlify(input []rune, length int) bool {
	if length == 0 {
		return false
	}

	lastIndex := length - 1
	lastBottomIndex := lastIndex

	// Calculates the last non-right whitespace index
	for ; lastIndex > 0 && input[lastIndex] == singleSpaceRune; lastIndex-- {
	}

	if lastIndex == 0 {
		return false
	}

	availableSpaces := lastBottomIndex - lastIndex

	// Counts Whitespaces
	whitespaceCount := 0
	for i := lastIndex; i >= 0; i-- {
		if input[i] == singleSpaceRune {
			whitespaceCount++
		}
	}

	// Checks if we have enough whitespaces to urlify the ones in between
	if whitespaceCount*2-availableSpaces > 0 {
		return false
	}

	for i, bottomIndex := lastIndex, length-1; i >= 0; i-- {
		if input[i] != singleSpaceRune {
			input[bottomIndex] = input[i]
			bottomIndex--
		} else {
			whitespaceCount++

			input[bottomIndex] = '0'
			input[bottomIndex-1] = '2'
			input[bottomIndex-2] = '%'

			bottomIndex = bottomIndex - 3
		}
	}

	return true
}
