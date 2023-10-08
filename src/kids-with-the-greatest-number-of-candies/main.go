package kids_with_the_greatest_number_of_candies

// topic 👉 https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var greatestAmount int
	for _, candyNumber := range candies {
		if candyNumber > greatestAmount {
			greatestAmount = candyNumber
		}
	}

	var result []bool
	for _, candyNumber := range candies {
		result = append(result, candyNumber+extraCandies >= greatestAmount)
	}

	return result
}
