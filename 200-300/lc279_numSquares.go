package main

import (
	"math"
	"sort"
)

var cell [101]int

func init() {
	for i := 0; i <= 100; i++ {
		cell[i] = i * i
	}
}

func numSquares(n int) int {
	// 本质上还是一个背包问题, 只不过求的是最小值

	var dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		var limit = sort.Search(101, func(k int) bool {
			return cell[k] > i
		}) - 1
		for j := 1; j <= limit; j++ {
			dp[i] = min(dp[i], dp[i-cell[j]]+1)
		}
	}

	return dp[n]
}

func numSquares2(n int) int {
	// 当DP写出来之后在沾沾自喜时, 大佬已经玩起了数学

	// 4平方和定理: 任何一个正整数都可以表示成不超过四个整数的平方之和

	isSquare := func(num int) bool {
		sqrt := int(math.Sqrt(float64(num)))
		return sqrt*sqrt == num
	}

	// 1. 本身就是平方数
	if isSquare(n) {
		return 1
	}

	// 2. 当 n=4^k*(8m+7) 时, 由于 4^k 可以提取出来, 所以结果也是 4
	{
		num := n
		for num&3 == 0 {
			num >>= 2
		}

		if num&7 == 7 {
			return 4
		}
	}

	// 3. i*i + j*j = n, 也就是说 n 可以由两个平方数组成
	for i := 1; i*i <= n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	// 4. 以上都不满足, 那就是 3 了
	return 3
}

func main() {
	// for i := 1; i <= 100; i++ {
	//	var limit = sort.Search(101, func(k int) bool {
	//		return cell[k] > i
	//	}) - 1
	//	fmt.Println(limit)
	// }
	for i := 1; i <= 100; i++ {
		println(i, numSquares(i))
	}
}
