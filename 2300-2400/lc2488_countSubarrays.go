package main

import "fmt"

func countSubarrays(nums []int, k int) int {
	kIndex := -1
	for i, num := range nums {
		if num == k {
			kIndex = i
			break
		}
	}
	ans := 0
	counts := map[int]int{}
	counts[0] = 1 // 只有k的情况, 至少一个
	sum := 0      // > k : 1, = k : 0, < k : -1
	// 按和k之间的大小关系计算出来的前缀和: 当前位置 大于k的数字的个数
	for i, num := range nums {
		sum += sign(num - k)
		if i < kIndex {
			// 只需要统计前半部分, 用来快速计算区间
			counts[sum]++
		} else {
			// 前缀和相同, 意味着 在 [left+1, right]这段区间内,

			// 包含 kIndex, 并且大于k和小于k的数字个数相同
			prev0 := counts[sum]
			// 前缀和-1, 意味着大于k的数比小于k的数少一个
			prev1 := counts[sum-1]
			fmt.Println(sum, prev0, prev1)
			ans += prev0 + prev1
		}
	}
	return ans
}

func sign(num int) int {
	if num == 0 {
		return 0
	}
	if num > 0 {
		return 1
	}
	return -1
}
