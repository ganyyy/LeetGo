package main

import "sort"

func maxIceCreamOld(costs []int, coins int) int {
	// 这不就是一个贪心?

	// 考察的是排序嘛?
	sort.Ints(costs)

	var i int
	for i < len(costs) {
		if costs[i] > coins {
			break
		}
		coins -= costs[i]
		i++
	}

	return i
}

func maxIceCream(costs []int, coins int) int {
	// 计数排序
	var cnt = [1e5 + 1]int32{}

	for _, c := range costs {
		cnt[c]++
	}

	var ret int
	for i := 1; i <= 1e5; i++ {
		if cnt[i] == 0 {
			continue
		}

		ret += min(int(cnt[i]), coins/i)
		coins -= int(cnt[i]) * i
		if coins <= i {
			break
		}
	}

	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
