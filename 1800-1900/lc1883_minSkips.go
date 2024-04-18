package main

import "math"

func minSkips(dist []int, speed int, hoursBefore int) int {
	n := len(dist)
	// f[i][j] 前i个节点休息j次所花的最短时间*speed
	// i ∈ [1, n], j ∈ [0, i]
	f := make([][]int64, n+1)
	for i := range f {
		f[i] = make([]int64, n+1)
		for j := range f[i] {
			f[i][j] = math.MaxInt64 / 2
		}
	}
	f[0][0] = 0
	// 这里的跳过, 指的是从 i-1 -> i(dist[i-1])
	// 路程时间是距离dist[i]/speed, 注意: 这里统一乘了speed避免了精度转换的问题!

	// 假设不跳过, 那么从i-1 -> i 需要的时间是路程时间+取整的等待时间
	// 假设跳过, 那么从i-1 -> i 所需要的时间就是路程自身的时间
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if j != i {
				// 不跳过, 相当于保留上一个点[i-1]的跳过次数[j]
				// 因为不跳过, 所以到达这个点的耗时等同于到达上一个点的事件+取整所需要补充的时间
				// 这里相当于按照speed进行了向上取整
				f[i][j] = min(f[i][j], ((f[i-1][j]+int64(dist[i-1])-1)/int64(speed)+1)*int64(speed))
			}
			if j != 0 {
				// 跳过, 相当于上一个节点[i-1]的
				f[i][j] = min(f[i][j], f[i-1][j-1]+int64(dist[i-1]))
			}
		}
	}
	totalCost := int64(hoursBefore) * int64(speed)
	for j := 0; j <= n; j++ {
		// 到达终点的最小休息次数
		if f[n][j] <= totalCost {
			return j
		}
	}
	return -1
}
