//go:build ignore

package main

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	// 无脑贪心(?)
	n, m := len(rowSum), len(colSum)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	i, j := 0, 0
	for i < n && j < m {
		// 取行列和中的最小值
		v := min(rowSum[i], colSum[j])
		matrix[i][j] = v
		rowSum[i] -= v
		colSum[j] -= v
		// 如果行/列的剩余值为0, 则表示需要换行了
		if rowSum[i] == 0 {
			i++
		}
		if colSum[j] == 0 {
			j++
		}
	}
	return matrix
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
