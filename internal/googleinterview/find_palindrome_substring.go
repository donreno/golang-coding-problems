package googleinterview

// FindAllPalindromeWords finds all palindrome words
func FindAllPalindromeWords(word string) []string {
	return findAllPalindromeWords(word, make(map[string]bool))
}

func findAllPalindromeWords(word string, memo map[string]bool) []string {
	if memo[word] {
		return []string{}
	}

	palindromes := []string{}
	wordAsRunes := []rune(word)

	for i := 1; i < len(wordAsRunes); i++ {
		firstWord := string(word[0:i])
		if len(firstWord) >= 2 && isPalindrome(firstWord) && !memo[firstWord] {
			memo[firstWord] = true
			palindromes = append(palindromes, firstWord)
		}

		secondWord := string(word[i:])
		if len(secondWord) >= 2 && isPalindrome(secondWord) && !memo[secondWord] {
			memo[secondWord] = true
			palindromes = append(palindromes, secondWord)
		}

		palindromes = append(palindromes, findAllPalindromeWords(string(word[i:]), memo)...)
	}

	return palindromes
}

func isPalindrome(word string) bool {
	wordAsRune := []rune(word)

	left, right := 0, len(word)-1

	for left < right {
		if wordAsRune[left] != wordAsRune[right] {
			return false
		}

		left++
		right--
	}

	return true
}
