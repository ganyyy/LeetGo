package main

import "sort"

func canReorderDoubled(arr []int) bool {
	cnt := make(map[int]int, len(arr))
	for _, x := range arr {
		cnt[x]++
	}
	if cnt[0]%2 == 1 {
		return false
	}

	val := make([]int, 0, len(cnt))
	for x := range cnt {
		val = append(val, x)
	}
	sort.Slice(val, func(i, j int) bool { return abs(val[i]) < abs(val[j]) })

	for _, x := range val {
		if cnt[2*x] < cnt[x] { // 无法找到足够的 2x 与 x 配对
			return false
		}
		// 这一步减完之后, 后来的2*x 计数就会成为0了
		cnt[2*x] -= cnt[x]
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
