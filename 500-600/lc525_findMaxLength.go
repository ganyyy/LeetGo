package main

import "fmt"

func findMaxLength(nums []int) int {
	// 451/564
	//var ret int
	//
	//// dp[i] 表示 到i点 最长的连续子数组长度
	//var dp = make([]int, len(nums))
	//
	//var check int
	//var sub int
	//for i := 1; i < len(nums); i++ {
	//	if nums[i] == 0 {
	//		check = 1
	//	} else {
	//		check = 0
	//	}
	//
	//	if dp[i-1] != 0 && i-dp[i-1]-sub >= 0 && nums[i-dp[i-1]-sub] == check {
	//		dp[i] += dp[i-1] + 2
	//	} else if nums[i-1] == check {
	//		sub = 1
	//		dp[i] = 2
	//	}
	//}
	//for i, v := range dp {
	//	if pre := i - v; pre >= 0 && pre < len(dp)-1 && dp[i-v+1] == 0 {
	//		// 如果能接上的话, 就更新一下最大的长度
	//		dp[i] += dp[i-v]
	//	}
	//	ret = max(ret, dp[i])
	//}
	//fmt.Println(dp)
	//return ret

	// 差分数组
	// 本质还是前缀和.

	var ret int
	for i, v := range nums {
		if v == 0 {
			nums[i] = -1
		}
	}

	var m = make(map[int]int)
	var sum int
	for i, v := range nums {
		sum += v
		// 等于0说明[:i] 0和1相等
		if sum == 0 && i > ret {
			ret = i + 1
		}
		// 如果前边存在同样的和, 说明 [idx:i]之间的0和1相等
		if idx, ok := m[sum]; ok {
			ret = max(ret, i-idx)
		} else {
			m[sum] = i
		}
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	findMaxLength([]int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0, 1, 1})
}
