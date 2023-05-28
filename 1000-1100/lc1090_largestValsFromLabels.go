package main

import "sort"

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	// 傻逼了, 脑子咋想的啊, 想什么聚合 label啊, 直接value排序不得了吗

	// 1. 按照 value 排序
	n := len(values)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		// 降序, 并且保证 idx 的顺序
		// 这个有意思啊, 学到了
		// 相当于借助value对idx进行排序
		return values[idx[i]] > values[idx[j]]
	})

	// 2. 按照 label 聚合
	ans, choose := 0, 0
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		label := labels[idx[i]]
		// 2.1 判断是否达到了 useLimit
		if cnt[label] == useLimit {
			continue
		}
		choose++
		ans += values[idx[i]]
		cnt[label]++
		// 2.2 判断是否达到了 numWanted
		if choose == numWanted {
			break
		}
	}
	return ans
}
