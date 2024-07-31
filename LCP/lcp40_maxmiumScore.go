package main

import "slices"

func maxmiumScore(cards []int, cnt int) int {
	slices.SortFunc(cards, func(a, b int) int { return b - a })
	s := 0
	for _, v := range cards[:cnt] {
		s += v
	}
	if s%2 == 0 { // s 是偶数
		return s
	}

	replaceSum := func(x int) int {
		for _, v := range cards[cnt:] {
			if v%2 != x%2 { // 找到一个最大的奇偶性和 x 不同的数
				return s - x + v // 用 v 替换 s
			}
		}
		return 0
	}
	// np
	// x 是前N个数字中的最小值, 再从后边找一个和x的奇偶不一样的最大值进行替换.
	// 然后就是从前边找一个和x的奇偶不一样的最小值, 再从后边找一个和这个值奇偶性不同的值进行替换
	x := cards[cnt-1]
	ans := replaceSum(x)            // 替换 x
	for i := cnt - 2; i >= 0; i-- { // 前 cnt-1 个数
		if cards[i]%2 != x%2 { // 找到一个最小的奇偶性和 x 不同的数
			ans = max(ans, replaceSum(cards[i])) // 替换
			break
		}
	}
	return ans
}
