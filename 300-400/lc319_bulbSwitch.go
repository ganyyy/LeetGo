package main

import "math"

func bulbSwitch(n int) int {
	// 因子个数为奇数
	return int(math.Sqrt(float64(n)))
}
