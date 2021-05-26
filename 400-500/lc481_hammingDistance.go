package main

func hammingDistance(x int, y int) int {
	var ret int

	//bits.OnesCount(uint(x ^ y))

	for v := x ^ y; v != 0; v = v & (v - 1) {
		ret++
	}
	return ret
}
