package unique_number_of_occurrences

// topic 👉 https://leetcode.com/problems/unique-number-of-occurrences/

func uniqueOccurrences(arr []int) bool {
	occurrencesMap := make(map[int]int)

	for index := range arr {
		occurrencesMap[arr[index]]++
	}

	seen := make(map[int]bool)
	for _, occurrence := range occurrencesMap {
		if seen[occurrence] {
			return false
		}
		seen[occurrence] = true
	}
	return true

}
