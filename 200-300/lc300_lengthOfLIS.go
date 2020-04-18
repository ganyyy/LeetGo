package main

import "fmt"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	var res, t int
	for i := 1; i < len(nums); i++ {
		t = 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				t = max(t, dp[j])
			}
		}
		dp[i] = t + 1
		res = max(dp[i], res)
	}
	return res
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
