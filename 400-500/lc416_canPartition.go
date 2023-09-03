package main

import "sort"

func canPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	// 只有是偶数的情况下才可能分为两部分
	if sum&1 != 0 {
		return false
	}
	var mid = sum >> 1
	// 草, 是一个背包问题.
	// 总的容量是 mid, 然后从 nums中选取任意数, 如果存在 dp[mid], 则可以认为是可以分割的
	// 一共存在 len(nums)个物品, 背包的容量是 mid
	// 01背包核心得需要倒叙遍历, 否则会出现重复计算的情况
	var dp = make([]bool, mid+1)
	for i := 0; i < len(nums); i++ {
		for s := mid; s >= nums[i]; s-- {
			if i == 0 {
				dp[s] = s == nums[i]
			} else {
				dp[s] = dp[s] || dp[s-nums[i]]
			}
		}
	}

	return dp[mid]
}

func canPartition2(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	sub := sum / 2

	// 当成一个背包问题来搞也不是不行
	// 这是01背包?
	// 要么选, 要么不选

	// 置换一下: 容量是sub, 价值都是1, 重量为nums[i]
	// 怎么避免重复选取呢..?

	sort.Ints(nums)
	var dp = make([]bool, sub+1)
	dp[0] = true
	for _, num := range nums {
		for i := sub; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
		if dp[sub] {
			return true
		}
	}
	return false
}

func main() {

}
