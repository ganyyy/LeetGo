package main

func findComplement(num int) int {
	var cur = 1

	for cur <= num {
		cur <<= 1
	}
	return cur - num - 1
}
