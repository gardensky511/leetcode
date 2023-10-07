package reverse_words_in_a_string_iii

import "strings"

// topic 👉 https://leetcode.com/problems/reverse-words-in-a-string-iii/

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	var reversedWords []string

	for _, word := range words {
		var reversedWord []rune

		for index := len(word) - 1; index >= 0; index-- {
			reversedWord = append(reversedWord, rune(word[index]))
		}

		reversedWords = append(reversedWords, string(reversedWord))
	}

	return strings.Join(reversedWords, " ")
}
