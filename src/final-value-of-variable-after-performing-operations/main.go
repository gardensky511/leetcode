package final_value_of_variable_after_performing_operations

import (
	"strings"
)

// topic 👉 https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

func finalValueAfterOperations1(operations []string) int {
	X := 0

	for _, operation := range operations {
		if strings.Contains(operation, "++") {
			X++
		} else {
			X--
		}
	}

	return X
}

func finalValueAfterOperations2(operations []string) int {
	X := 0

	values := map[byte]int{'+': 1, '-': -1}

	for _, operation := range operations {
		X += values[operation[1]]
	}

	return X
}
