package main

import "math"

func nthSuperUglyNumber(n int, primes []int) int {
	var cnt = make([]int, len(primes))
	var tmp = make([]int, 0, n)

	tmp = append(tmp, 1)

	for i := 2; i <= n; i++ {
		// 各个数字入队
		var m = math.MaxInt32
		for j, v := range primes {
			if t := v * tmp[cnt[j]]; t < m {
				m = t
			}
		}
		tmp = append(tmp, m)

		// 增加对应位置的索引
		for j, v := range primes {
			if v*tmp[cnt[j]] == m {
				cnt[j]++
			}
		}
	}
	return tmp[len(tmp)-1]
}
