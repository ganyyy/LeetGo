package main

import "fmt"

func jump(nums []int) int {
	// 指定位置的数相加 >= len-1
	// 记录选取的数的最小个数
	// 不需要处理最后一个数
	// 状态转移方程: dp[i][j] 从 i -> j 最少要几步
	ln := len(nums)
	if ln == 1 {
		return 0
	}
	dp := make([]int, ln)
	for i := 0; i < ln-1; i++ {
		val := dp[i]
		if i+nums[i] >= ln-1 {
			return val + 1
		}
		for j := i + 1; j <= i+nums[i]; j++ {
			if dp[j] == 0 {
				dp[j] = val + 1
			} else {
				dp[j] = min(dp[j], val+1)
			}
		}
	}
	return dp[ln-1]
}

// 上贪心算法
func jump2(nums []int) int {
	end, maxPos, steps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		// 更新能跳到的最远位置
		maxPos = max(maxPos, i+nums[i])
		if i == end {
			// 到了最远位置就更新一下下一次的终点, 并加一步
			end = maxPos
			steps++
		}
	}
	return steps
}

func main() {
	fmt.Println(jump([]int{1, 2, 3, 4, 5}))
}
