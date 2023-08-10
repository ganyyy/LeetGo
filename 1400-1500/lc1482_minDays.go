package main

import (
	"math"
	"sort"
)

func minDays(bloomDay []int, m int, k int) int {
	// emm, 也是值二分法

	var mi, mx = math.MaxInt32, math.MinInt32
	for _, v := range bloomDay {
		mi = min(mi, v)
		mx = max(mx, v)
	}

	// 需要建立快速索引吗..?

	var ret = sort.Search(mx-mi+1, func(i int) bool {
		// 判断能不能找到可以满足 m束花, 每束k朵

		// 当前对应的天数
		var cur = i + mi

		var mc, mk int
		for _, v := range bloomDay {
			if v > cur {
				mk = 0
				continue
			}
			mk++
			if mk >= k {
				mc++
				mk = 0
			}
			if mc >= m {
				return true
			}
		}
		return false
	})

	if ret+mi > mx {
		return -1
	}
	return ret + mi
}
