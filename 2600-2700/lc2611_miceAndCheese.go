package main

import "sort"

func miceAndCheese(reward1, reward2 []int, k int) (ans int) {
	for i, x := range reward2 {
		ans += x // 先全部给第二只老鼠
		reward1[i] -= x
	}
	// 选差值最大的前k个, 给🐀1, 那么整体的和肯定就是最大的
	sort.Sort(sort.Reverse(sort.IntSlice(reward1)))
	for _, x := range reward1[:k] {
		ans += x
	}
	return
}
