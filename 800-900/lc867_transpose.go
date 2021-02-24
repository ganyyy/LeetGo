package main

func transpose(matrix [][]int) [][]int {
	var n, m = len(matrix), len(matrix[0])
	var ret = make([][]int, m)
	for i := 0; i < m; i++ {
		ret[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ret[i][j] = matrix[j][i]
		}
	}
	return ret
}
