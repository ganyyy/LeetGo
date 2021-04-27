package main

import "math"

func judgeSquareSum(c int) bool {
	/*

	   x^2 + y^2 = c

	   (x-y)^2 = c - 2xy

	   1. c - 2xy 必须是一个整数的平方
	   2. 直接对 c 开根号, c - 结果的平方就是
	*/

	// 双指针还能真么用吗.. 大意了, 没有闪
	var left, right = 0, sqrt(c)

	for left <= right {
		if t := left*left + right*right; t > c {
			right--
		} else if t < c {
			left++
		} else {
			return true
		}
	}
	return false
}

func sqrt(a int) int {
	return int(math.Sqrt(float64(a)))
}
