package main

func xorGame(nums []int) bool {
	if len(nums)&1 == 0 {
		return true
	}
	var sum int

	for _, v := range nums {
		sum ^= v
	}

	return sum == 0
}
