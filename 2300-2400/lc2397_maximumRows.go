package main

import "math/bits"

func maximumRows(mat [][]int, numSelect int) (ans int) {
	m, n := len(mat), len(mat[0])
	// mask: 每一行上选取的位置
	mask := make([]int, m)
	for i, row := range mat {
		for j, x := range row {
			mask[i] |= x << j
		}
	}

	for subset := 0; subset < 1<<n; subset++ {
		if bits.OnesCount(uint(subset)) != numSelect {
			continue
		}
		coveredRows := 0
		for _, row := range mask {
			// 如果当前选取的subset恰好包括了当前行所有的1, 说明清空之后整体就变成了0
			if row&subset == row {
				coveredRows++
			}
		}
		ans = max(ans, coveredRows)
	}
	return
}
