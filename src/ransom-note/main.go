package ransom_note

import "strings"

// topic 👉 https://leetcode.com/problems/ransom-note/

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	rest := magazine

	for _, letter := range ransomNote {
		stringLetter := string(letter)

		if strings.Contains(rest, stringLetter) {
			rest = strings.Replace(rest, stringLetter, "", 1)
		} else {
			return false
		}
	}

	return true
}
