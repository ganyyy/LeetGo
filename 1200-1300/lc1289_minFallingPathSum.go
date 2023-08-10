package main

import "math"

func minFallingPathSum(grid [][]int) int {
	n := len(grid)
	// 累加到上一行的最小值和次小值
	firstMinSum, secondMinSum := 0, 0
	// 上一行的最小值下标
	firstMinIndex := -1

	for i := 0; i < n; i++ {
		// 当前行的最小值, 次小值, 和最小值对应的下标
		curFirstMinSum, curSecondMinSum := math.MaxInt, math.MaxInt
		curFirstMinIndex := -1

		for j := 0; j < n; j++ {
			sum := grid[i][j]
			if j != firstMinIndex {
				sum += firstMinSum
			} else {
				sum += secondMinSum
			}
			if sum < curFirstMinSum {
				curSecondMinSum, curFirstMinSum = curFirstMinSum, sum
				curFirstMinIndex = j
			} else if sum < curSecondMinSum {
				curSecondMinSum = sum
			}
		}
		firstMinSum, secondMinSum = curFirstMinSum, curSecondMinSum
		firstMinIndex = curFirstMinIndex
	}
	return firstMinSum
}
