package main

func isToeplitzMatrix(matrix [][]int) bool {
	var n, m = len(matrix), len(matrix[0])

	var check = func(x, y int) bool {
		var val = matrix[x][y]
		for x, y = x+1, y+1; x < n && y < m; x, y = x+1, y+1 {
			if val != matrix[x][y] {
				return false
			}
		}
		return true
	}

	// 中线
	if !check(0, 0) {
		return false
	}

	// 竖着来
	for i := 1; i < n; i++ {
		if !check(i, 0) {
			return false
		}
	}
	// 横着来
	for i := 1; i < m; i++ {
		if !check(0, i) {
			return false
		}
	}

	return true
}
