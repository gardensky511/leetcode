package rank_transform_of_an_array

import "sort"

// topic 👉 https://leetcode.com/problems/rank-transform-of-an-array/

func arrayRankTransform(arr []int) []int {
	ranks := make([]int, len(arr))

	sortedArr := removeDuplicates(arr)
	sort.Ints(sortedArr)

	for index, element := range arr {
		rank := sort.SearchInts(sortedArr, element) + 1
		ranks[index] = rank
	}

	return ranks
}

func removeDuplicates(arr []int) []int {
	uniqueMap := make(map[int]bool)
	var uniqueArr []int

	for _, item := range arr {
		if !uniqueMap[item] {
			uniqueArr = append(uniqueArr, item)
			uniqueMap[item] = true
		}
	}

	return uniqueArr
}
