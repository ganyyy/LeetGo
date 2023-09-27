package main

import "math/bits"

func hammingWeight(num uint32) int {
	var count int
	// 核心想法是通过-1将低位的1变成0
	// 如 0100 0011
	for num > 0 {
		num &= num - 1
		count++
	}
	bits.OnesCount32(num)
	return count
}

func hammingWeight2(n uint32) int {
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	n = (n & 0x0f0f0f0f) + ((n >> 4) & 0x0f0f0f0f)
	n = (n & 0x00ff00ff) + ((n >> 8) & 0x00ff00ff)
	n = (n & 0x0000ffff) + ((n >> 16) & 0x0000ffff)
	return int(n)
}

func main() {

}
