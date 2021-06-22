package main

func hammingWeight(num uint32) (cnt int) {
	//return bits.OnesCount(uint(num))
	for num != 0 {
		num &= num - 1
		cnt++
	}
	return
}
