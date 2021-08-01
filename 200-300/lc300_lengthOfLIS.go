package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS(nums []int) int {
	// 先写一个O^2级别的
	var dp = make([]int, len(nums))
	dp[0] = 1
	var ret = 1

	// dp[i]表示nums[:i]中的最长递增子序列

	for i := 1; i < len(nums); i++ {
		var t = 0
		for j := 0; j < i; j++ {
			// 从头开始, 找到所有比i小的数字, 计算最大值
			if nums[j] < nums[i] {
				t = max(t, dp[j])
			}
		}
		// t 表示比nums[i]小的最长子序列, +1表示将dp[i]加入到队列中
		dp[i] = t + 1
		ret = max(ret, dp[i])
	}
	return ret
}

func lengthOfLIS2(nums []int) int {
	// 正儿八经的 LIS

	// dp整个数组都是有序的
	// dp中存在的数字不一定是正确的结果即不可回溯
	// 但是长度一定和最终解长度相等
	var dp []int

	for _, v := range nums {
		// 二分查找该值在dp数组中的位置
		// 三种结果
		// 1. idx < len(dp), 此时可能是存在该值, 或者不存在该值, 但是小于数组中的最大值
		// 那么将dp[idx]上的数值进行替换, 可以保证当前的dp是满足单调递增条件
		// 2. idx == len(dp), 加入该值, 数组仍然保持递增
		var idx = sort.SearchInts(dp, v)
		if idx < len(dp) {
			dp[idx] = v
		} else {
			dp = append(dp, v)
		}
	}

	return len(dp)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
