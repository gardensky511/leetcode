package concatenation_of_array

// topic 👉 https://leetcode.com/problems/concatenation-of-array/description/

func getConcatenationFirstSolution(nums []int) []int {
	slice := nums[:]
	return append(slice, slice...)
}

func getConcatenationSecondSolution(nums []int) []int {
	return append(nums, nums...)
}
