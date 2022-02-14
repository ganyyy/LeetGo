package main

import "math"

func luckyNumbers(matrix [][]int) []int {
	var m, n = len(matrix), len(matrix[0])

	var minM = make([]int, m) // 每一行的最小值
	var maxN = make([]int, n) // 每一列的最大值

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < m; i++ {
		minM[i] = math.MaxInt32
		for j := 0; j < n; j++ {
			var v = matrix[i][j]
			minM[i] = min(minM[i], v)
			maxN[j] = max(maxN[j], v)
		}
	}

	// 整合答案
	if m > n {
		minM, maxN = maxN, minM
	}

	var check func(v int) bool

	var minMap = make(map[int]bool, m)
	for _, v := range minM {
		minMap[v] = true
	}
	check = func(v int) bool {
		return minMap[v]
	}

	var ret = make([]int, 0, n>>2)

	for _, v := range maxN {
		if !check(v) {
			continue
		}
		ret = append(ret, v)
	}

	return ret
}
