package number_of_steps_to_reduce_a_number_to_zero

// topic 👉 https://leetcode.com/problems/number-of-steps-to-reduce-a-number-to-zero/

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}

	var steps int

	rest := num
	for {
		if rest%2 == 0 {
			rest = rest / 2
		} else {
			rest = rest - 1
		}
		steps += 1

		if rest == 0 {
			break
		}
	}

	return steps
}
