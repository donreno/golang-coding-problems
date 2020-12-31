package golang_coding_problems

var singleSpaceRune = rune(' ')
var urlSpace = []rune("%20")

// Urlify replaces non-left or right whitespaces with '%20' on given url
func Urlify(input []rune, length int) bool {
	if length == 0 {
		return false
	}

	lastIndex := length - 1

	for ; lastIndex > 0 && input[lastIndex] == singleSpaceRune; lastIndex-- {
	}

	if lastIndex == 0 {
		return false
	}

	tmpUrlified := []rune("")

	for i := 0; i <= lastIndex; i++ {
		if input[i] != singleSpaceRune {
			tmpUrlified = append(tmpUrlified, input[i])
		} else {
			tmpUrlified = append(tmpUrlified, urlSpace...)
		}
	}

	if len(tmpUrlified) != length {
		return false
	}

	for i, r := range tmpUrlified {
		input[i] = r
	}

	_ = tmpUrlified

	return true
}
