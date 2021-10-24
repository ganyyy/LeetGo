package main

func searchMatrix(matrix [][]int, target int) bool {

	// 从左下角开始找
	var col = 0
	var row = len(matrix) - 1

	for row >= 0 && col < len(matrix[0]) {
		if target == matrix[row][col] {
			return true
		} else if target < matrix[row][col] {
			row--
		} else {
			col++
		}
	}

	return false
}
