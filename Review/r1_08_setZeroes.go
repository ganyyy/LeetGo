package main

func setZeroes(matrix [][]int) {
	row := len(matrix)
	if row == 0 {
		return
	}
	col := len(matrix[0])
	if col == 0 {
		return
	}

	// 判断首行/首列是否存在0
	var zRow, zCol bool
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			zCol = true
		}
	}
	for i := 0; i < col; i++ {
		if matrix[0][i] == 0 {
			zRow = true
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < row; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < col; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < col; i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < row; j++ {
				matrix[j][i] = 0
			}
		}
	}

	if zRow {
		for i := 0; i < col; i++ {
			matrix[0][i] = 0
		}
	}
	if zCol {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}
