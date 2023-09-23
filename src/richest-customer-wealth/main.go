package richest_customer_wealth

// topic 👉 https://leetcode.com/problems/richest-customer-wealth/

func maximumWealth(accounts [][]int) int {
	var maxWealth int

	for _, account := range accounts {
		var sum int

		for _, value := range account {
			sum += value
		}

		if sum > maxWealth {
			maxWealth = sum
		}
	}

	return maxWealth
}
