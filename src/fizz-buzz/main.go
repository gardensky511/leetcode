package fizz_buzz

// topic 👉 https://leetcode.com/problems/fizz-buzz/

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	answer := make([]string, n)

	for index := range answer {
		order := index + 1

		if order%15 == 0 {
			answer[index] = "FizzBuzz"
			continue
		}
		if order%3 == 0 {
			answer[index] = "Fizz"
			continue
		}
		if order%5 == 0 {
			answer[index] = "Buzz"
			continue
		}

		answer[index] = strconv.Itoa(order)
	}

	return answer
}
