package main

import (
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {
	m := len(n)
	candidates := []int{int(math.Pow10(m-1)) - 1, int(math.Pow10(m)) + 1} // 9999, 10001

	// 12345 -> 123
	//  5678 -> 56
	selfPrefix, _ := strconv.Atoi(n[:(m+1)/2])                            // 前半部分
	for _, x := range []int{selfPrefix - 1, selfPrefix, selfPrefix + 1} { // 122, 123, 124
		y := x
		// 123 -> 12
		// 56  -> 56
		if m&1 == 1 {
			y /= 10
		}

		// 123 -> 12321
		//  56 -> 5665
		for ; y > 0; y /= 10 {
			x = x*10 + y%10
		}
		candidates = append(candidates, x)
	}

	ans := -1
	selfNumber, _ := strconv.Atoi(n)
	for _, candidate := range candidates {
		// 迭代所有的备选数字, 选出差值最小的回文数
		if candidate != selfNumber {
			if ans == -1 ||
				abs(candidate-selfNumber) < abs(ans-selfNumber) ||
				abs(candidate-selfNumber) == abs(ans-selfNumber) && candidate < ans {
				ans = candidate
			}
		}
	}
	return strconv.Itoa(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
