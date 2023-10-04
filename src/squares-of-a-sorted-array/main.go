package squares_of_a_sorted_array

import "sort"

// topic 👉 https://leetcode.com/problems/squares-of-a-sorted-array/

func sortedSquares(nums []int) []int {
	squared := make([]int, len(nums))

	for index := range nums {
		squared[index] = nums[index] * nums[index]
	}
	sort.Ints(squared)

	return squared
}
