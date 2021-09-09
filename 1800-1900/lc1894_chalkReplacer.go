package main

func chalkReplacer(chalk []int, k int) int {
	// 先求和, 再取余

	var sum int
	for _, v := range chalk {
		sum += v
	}

	k %= sum

	for i, v := range chalk {
		k -= v
		if k < 0 {
			return i
		}
	}

	return 0
}
