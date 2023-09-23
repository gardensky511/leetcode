package shuffle_the_array

// topic 👉 https://leetcode.com/problems/shuffle-the-array/

func shuffle(nums []int, n int) []int {
	var answer []int

	for i := 0; i < n; i++ {
		answer = append(answer, nums[i], nums[i+n])
	}

	return answer
}
