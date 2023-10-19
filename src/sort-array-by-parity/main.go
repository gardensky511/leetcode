package sort_array_by_parity

// topic 👉 https://leetcode.com/problems/sort-array-by-parity/

func sortArrayByParity(nums []int) []int {
	var answer []int

	for _, num := range nums {
		if num%2 == 0 {
			answer = append([]int{num}, answer...)
		} else {
			answer = append(answer, num)
		}
	}

	return answer
}
