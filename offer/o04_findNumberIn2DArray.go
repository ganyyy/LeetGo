package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	var m = len(matrix)
	if m == 0 {
		return false
	}
	var n = len(matrix[0])
	if n == 0 {
		return false
	}

	var i, j = 0, m - 1

	for i < n && j >= 0 {
		var cur = matrix[j][i]
		if cur == target {
			return true
		}
		if cur > target {
			j--
		} else {
			i++
		}
	}

	return false

}
