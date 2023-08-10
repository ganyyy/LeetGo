package main

import "fmt"

func maxCoins(nums []int) int {
	ln := len(nums)

	if ln == 0 {
		return 0
	}
	if ln == 1 {
		return nums[0]
	}
	if ln == 2 {
		return nums[0]*nums[1] + max(nums[0], nums[1])
	}

	// 将开头和结尾填充进去
	allNum := make([]int, ln+2)
	allNum[0] = 1
	allNum[ln+1] = 1
	copy(allNum[1:], nums)

	// 选取可能的最大值
	// dp[i][j] 表示从 i-j 通过戳气球可以获取到的最大值(不包括i,j)
	dp := make([][]int, ln+2)
	for i := 0; i <= ln; i++ {
		dp[i] = make([]int, ln+2)
	}

	// i 的取值范围是 [0, ln], 等同于 从
	for i := ln; i >= 0; i-- {
		// j 的取值范围是 [1, ln+1]
		for j := i + 1; j < ln+2; j++ {
			// k 的取值范围是 [i+1,j-1]
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j],
					// 戳破 k 位置的气球, 求的的和
					dp[i][k]+dp[k][j]+allNum[i]*allNum[j]*allNum[k],
				)
			}
		}
	}

	return dp[0][ln+1]
}

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}
