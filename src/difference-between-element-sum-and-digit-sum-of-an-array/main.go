package difference_between_element_sum_and_digit_sum_of_an_array

// topic 👉 https://leetcode.com/problems/difference-between-element-sum-and-digit-sum-of-an-array/

func differenceOfSum(nums []int) int {
	var elementSum, digitSum int

	for _, num := range nums {
		elementSum += num

		for num != 0 {
			digitSum += num % 10
			num /= 10
		}
	}

	return elementSum - digitSum
}
