package main

import (
	"fmt"
	"math"
)

func superEggDrop2(K int, N int) int {
	//  当前有 k 个鸡蛋，可以尝试扔 m 次鸡蛋, 最坏情况下最多能确切测试一栋 n 层的楼
	//  比如说 dp[1][7] = 7 表示：
	//  现在有 1 个鸡蛋，允许你扔 7 次; 这个状态下最多给你 7 层楼，使得你可以确定楼层 F 使得鸡蛋恰好摔不碎
	//  (一层一层线性探查嘛)
	dp := make([][]int, K+1)
	for i := 0; i < K+1; i++ {
		dp[i] = make([]int, N+1)
	}
	var m int
	for dp[K][m] < N {
		m++
		for k := 1; k <= K; k++ {
			dp[k][m] = dp[k][m-1] + dp[k-1][m-1] + 1
		}
	}
	return m
}

func superEggDrop3(K, N int) int {
	// 将上述过程优化为一维数组
	dp := make([]int, K+1)
	var m int
	for dp[K] < N {
		m++
		for k := K; k >= 1; k-- {
			dp[k] += dp[k-1] + 1
		}
	}
	return m
}

func superEggDrop(K int, N int) int {
	// 简易的动态规划
	var dp func(int, int) int

	r := make(map[[2]int]int)

	// k : 鸡蛋数量
	// n : 楼层高度
	dp = func(k int, n int) int {
		// 一个鸡蛋最多需要N次
		if k == 1 {
			return n
		}
		// 0层楼不需要进行判断
		if n == 0 {
			return 0
		}
		if v, ok := r[[2]int{k, n}]; ok {
			return v
		}
		mx := math.MaxInt32
		for i := 1; i <= n; i++ {
			// 如果鸡蛋没碎, 就从 n-i 中找
			// 如果鸡蛋碎了, 就从 i-1 中找, 同时鸡蛋数量 k-1
			mx = min(mx, max(dp(k, n-i), dp(k-1, i-1))+1)
		}
		r[[2]int{k, n}] = mx
		return mx
	}

	return dp(K, N)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(superEggDrop2(3, 14))
}
