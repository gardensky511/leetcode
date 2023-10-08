package find_target_indices_after_sorting_array

import "sort"

// topic 👉 https://leetcode.com/problems/find-target-indices-after-sorting-array/

func targetIndices(nums []int, target int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	var indices []int
	for index := range sortedNums {
		if sortedNums[index] == target {
			indices = append(indices, index)
		}
	}

	return indices
}
