package build_array_from_permutation

// topic 👉 https://leetcode.com/problems/build-array-from-permutation/

func buildArray1(nums []int) []int {
	var ans []int

	for i := 0; i < len(nums); i++ {
		ans = append(ans, nums[nums[i]])
	}

	return ans
}

func buildArray2(nums []int) []int {
	ans := make([]int, len(nums))

	for index, num := range nums {
		ans[index] = nums[num]
	}

	return ans
}
