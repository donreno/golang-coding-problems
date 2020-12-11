package golang_coding_problems

import "strings"

const closeParentheses = "}])>"

var parenthesesMatches map[string]string

// CheckBalancedString checks if a string's parentheses are balanced
func CheckBalancedString(str string) bool {
	tokens := strings.Split(str, "")
	parenthesesStack := MakeStack()

	for _, token := range tokens {
		_, isOpenParentheses := parenthesesMatches[token]
		isCloseParentheses := strings.Contains(closeParentheses, token)

		if isOpenParentheses {
			parenthesesStack.Push(token)
		} else if isCloseParentheses {
			if parenthesesStack.Empty() {
				return false
			}

			lastOpenParentheses := parenthesesStack.Pop().(string)

			if parenthesesMatches[lastOpenParentheses] != token {
				return false
			}
		}
	}

	return true
}

func init() {
	parenthesesMatches = make(map[string]string)

	parenthesesMatches["["] = "]"
	parenthesesMatches["{"] = "}"
	parenthesesMatches["("] = ")"
	parenthesesMatches["<"] = ">"
}
