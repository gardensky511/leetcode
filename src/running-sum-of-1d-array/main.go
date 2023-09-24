package running_sum_of_1d_array

// topic 👉 https://leetcode.com/problems/running-sum-of-1d-array/

func runningSum(nums []int) []int {
	answer := make([]int, len(nums))

	var sum int
	for index, num := range nums {
		sum += num
		answer[index] = sum
	}

	return answer
}
