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
	isSquare := func(num int) bool {
		sqrt := int(math.Sqrt(float64(num)))
		return sqrt*sqrt == num
	}

	num := n
	for num&3 == 0 {
		num >>= 2
	}

	if num&7 == 7 {
		return 4
	}

	if isSquare(n) {
		return 1
	}

	for i := 1; i*i <= n; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	return 3
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//for i := 1; i <= 100; i++ {
	//	var limit = sort.Search(101, func(k int) bool {
	//		return cell[k] > i
	//	}) - 1
	//	fmt.Println(limit)
	//}
	for i := 1; i <= 100; i++ {
		println(i, numSquares(i))
	}
}
