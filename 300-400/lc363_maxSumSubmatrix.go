//go:build ignore

package main

import "math"

func maxSumSubMatrix(matrix [][]int, k int) int {
	m, n, maxArea := len(matrix), len(matrix[0]), math.MinInt64
	// O(cols ^ 2 * rows)

	// 以列为单位, 进行降维.
	for l := 0; l < n; l++ { // 枚举左边界
		rowSum := make([]int, m) // 左边界改变才算区域的重新开始
		for r := l; r < n; r++ { // 枚举右边界
			// 求当前列区间内, 每一行的和
			for i := 0; i < m; i++ { // 按每一行累计到 dp
				rowSum[i] += matrix[i][r]
			}
			maxArea = max(maxArea, dpMax(rowSum, k))
			if maxArea == k {
				return k // 尽量提前
			}
		}
	}

	return maxArea
}

// 在数组 arr 中，求不超过 k 的最大值
// 核心就是求不大于K的最小值
// 这个算法目前来看只能枚举处理
func dpMax(arr []int, k int) int {
	rollSum := arr[0]
	rollMax := rollSum

	// O(rows)
	// 先看所有行
	// 求出连续的最大值
	for i := 1; i < len(arr); i++ {
		if rollSum > 0 {
			rollSum += arr[i]
		} else {
			rollSum = arr[i]
		}

		rollMax = max(rollMax, rollSum)
	}
	// 如果连续的最大值都小于等于K..
	// 说真的, 这个不好说啊. 是否有必要提前返回?
	if rollMax <= k {
		return rollMax
	}
	// O(rows ^ 2)
	// 再看行区间的
	// 这就一暴力解法啊... 头疼

	// 这里确定相应的行区间的和
	maxArea := math.MinInt64
	for l := 0; l < len(arr); l++ {
		sum := 0
		for r := l; r < len(arr); r++ {
			sum += arr[r]
			if sum > maxArea && sum <= k {
				maxArea = sum
			}
			if maxArea == k {
				return k // 尽量提前
			}
		}
	}

	return maxArea
}
