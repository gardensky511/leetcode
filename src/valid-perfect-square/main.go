package valid_perfect_square

// topic 👉 https://leetcode.com/problems/valid-perfect-square/

import "math"

func isPerfectSquare(num int) bool {
	answer := math.Sqrt(float64(num))
	return answer == math.Trunc(answer)
}
