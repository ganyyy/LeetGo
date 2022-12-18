package main

import "math"

func minMoves(nums []int, k int) int {
	zero := countZero(nums)
	// 计算窗口
	if k-1 > len(zero) {
		// 无法完成的交换
		return math.MaxInt32
	}

	var cost int
	var left int
	var right = k - 2

	for i := left; i <= right; i++ {
		// 向左还是向右
		cost += zero[i] * (min(i+1, right-i+1))
	}

	pre := preSum(zero)
	getSum := func(l, r int) int {
		return pre[r+1] - pre[l]
	}
	minCost := cost
	add := k % 2
	for i, j := 1, k-1; j < len(zero); i, j = i+1, j+1 {
		mid := (i + j) / 2
		cost -= getSum(i-1, mid-1)
		cost += getSum(mid+add, j)
		minCost = min(minCost, cost)
	}
	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countZero(nums []int) []int {
	// 去掉首尾后, 连续0的个数
	ln := len(nums)
	var ret []int
	var i int
	for ; i < ln && nums[i] == 0; i++ {
	}
	for i < ln {
		j := i + 1
		for ; j < ln && nums[j] == 0; j++ {
		}
		if j < ln {
			ret = append(ret, j-i-1)
		}
		i = j
	}
	return ret
}

func preSum(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}
	// 填充一个开头0
	var ret = make([]int, 1, len(nums)+1)
	for i, v := range nums {
		ret = append(ret, v+ret[i])
	}
	return ret
}
