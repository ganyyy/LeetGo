package main

import (
	"container/heap"
	"sort"
)

func minRefuelStops(target, startFuel int, stations [][]int) (ans int) {
	// 将所有途径的油桶都带上, 如果当前剩余的油量不足, 就从过往带上的油桶中获取可使用的最大值
	fuel, prev, h := startFuel, 0, hp{}
	// station[0]: 加油站所处的位置
	// station[1]: 加油站能加的油的数量
	for i, n := 0, len(stations); i <= n; i++ {
		curr := target
		if i < n {
			curr = stations[i][0]
		}
		// 首先判断一下, 当前剩余的油量能不到到达第i个加油站
		fuel -= curr - prev
		for fuel < 0 && h.Len() > 0 {
			// 如果不能, 就从备用的油库中取出油量最大的那个
			// 直到当前油量>=0或者没有备用油库
			fuel += heap.Pop(&h).(int)
			// 每次取出一个都要+1
			ans++
		}
		if fuel < 0 {
			// 如果还是没有剩余油量, 直接说明无法到达
			return -1
		}
		if i < n {
			// 将当前加油站的剩余油量添加到备库中
			heap.Push(&h, stations[i][1])
			prev = curr
		}
	}
	return
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func minRefuelStops2(target, startFuel int, stations [][]int) int {
	n := len(stations)
	// dp[i] 表示加油i次可以到达的最远距离
	// 因为dp[i]依赖于dp[i-1], 所以需要从前向后迭代
	// 本质上还是贪心
	dp := make([]int, n+1)
	dp[0] = startFuel
	for i, s := range stations {
		for j := i; j >= 0; j-- {
			if dp[j] >= s[0] {
				dp[j+1] = max(dp[j+1], dp[j]+s[1])
			}
		}
	}
	for i, v := range dp {
		if v >= target {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
