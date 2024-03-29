package main

import "slices"

func minimumAddedCoins(coins []int, target int) (ans int) {
	slices.Sort(coins)
	// 假设现在已经满足 [0, s-1]区间内的所有整数
	s, i := 1, 0
	for s <= target {
		if i < len(coins) && coins[i] <= s {
			// x <= s, 可以直接添加x, 此时区间变成了 [0, s+x-1]
			s += coins[i]
			i++
		} else {
			// 此时集合中最大数就是s, 直接添加即可, 添加后的范围为[0, 2*s-1]
			// 因为得需要保证连续, 所以要想能到达的位置更远, 只能添加s
			s *= 2 // 必须添加 s
			ans++
		}
	}
	return
}
