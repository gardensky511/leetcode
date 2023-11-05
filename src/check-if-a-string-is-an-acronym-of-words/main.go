package check_if_a_string_is_an_acronym_of_words

// topic 👉 https://leetcode.com/problems/check-if-a-string-is-an-acronym-of-words/

func isAcronym(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	for i := 0; i < len(words); i++ {
		if words[i][0] != s[i] {
			return false
		}
	}

	return true
}
