package main

func isIdealPermutation(nums []int) bool {
	for i, x := range nums {
		if abs(x-i) > 1 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
