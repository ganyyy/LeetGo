package main

import "fmt"

func waysToChange(n int) int {
	a := [4]int{1, 5, 10, 25}
	dp := make([]int, n+1)
	mod := 1000000007
	dp[0] = 1 // 如果钱正好和a[i]相等时, 这相当于一种取法

	// 先从硬币开始取, 避免不同顺序的相同硬币算是不同解法的情况
	for _, i := range a {
		for j := 1; j <= n; j++ {
			if i <= j {
				dp[j] = (dp[j] + dp[j-i]) % mod
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(waysToChange(15))
}
