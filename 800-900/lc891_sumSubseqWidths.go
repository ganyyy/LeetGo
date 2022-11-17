package main

import "sort"

const MOD int = 1e9 + 7
const N int = 1e5

var pow = func() []int {
	ret := make([]int, N)
	ret[0] = 1
	for i := 1; i < N; i++ {
		ret[i] = ret[i-1] * 2 % MOD
	}
	return ret
}()

func sumSubseqWidths(nums []int) int {
	// 子序列和数据出现的位置无关
	sort.Ints(nums)

	// 对于任意一个数(下标>0), 均可以作为子数组的最大值和最小值
	// 当(i)作为最大值时, 其需要计算 2^i次       (加)
	// 当(i)作为最小值时, 其需要计算 2^(n-1-i)次 (减)
	var ret int
	n := len(nums)
	for i, v := range nums {
		cnt := (pow[i] - pow[n-i-1]) * v % MOD
		ret = (ret + cnt) % MOD
	}
	return ret
}
