package main

import (
	"math"
	"sort"
)

func maxPower(stations []int, r int, k int) int64 {
	length := len(stations)
	// 前缀和. preSum[i] = sum(stations[:i])
	preSum := make([]int, length+1)
	for idx, power := range stations {
		preSum[idx+1] = preSum[idx] + power
	}

	// 所有城市中最小的电量
	var minCityPower = math.MaxInt
	// 通过范围计算每个城市的实际电量
	for idx := range stations {
		// ...r-i...i...i+r...
		stations[idx] = preSum[min(idx+r+1, length)] - preSum[max(idx-r, 0)]
		minCityPower = min(minCityPower, stations[idx])
	}

	// 差分数组. 表示增加的额外的电站的影响范围
	var diff = make([]int, length)
	// 二分确认答案
	// k表示新增的电站的数量. 也就是说: k是增加的上限
	return int64(minCityPower + sort.Search(k, func(minPower int) bool {
		clear(diff)
		// minPower的范围是[0, k)
		// 我们实际计算的范围是[1, k], 所以得需要+1(?)
		minPower += minCityPower + 1
		// totalDiff 差分数组的累加和. 因为是区间累加, 所以是针对差分数组进行的计算
		// need      累计需要增加的电站的数量
		var totalDiff, need int
		// ...i...i+r...i+2*r...
		for idx, power := range stations {
			totalDiff += diff[idx]
			// add 表示这个位置需要增加的电站的数量
			add := minPower - power - totalDiff
			if add > 0 {
				// 这个位置需要增加一些电站
				need += add
				if need > k {
					// 此时意味着 minPower 太大了, 需要减少
					return true
				}
				totalDiff += add
				if end := idx + 2*r + 1; end < length {
					diff[end] -= add
				}
			}
		}
		// 说明 minPower 还可以增大(?)
		return false
	}))

}
