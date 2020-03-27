package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	ln, lm := len(matrix), len(matrix[0])

	row, col := 0, lm-1
	// 从右上角开始, 大的在下边, 小的在右边
	for row < ln && col > -1 {
		if matrix[row][col] > target {
			col--
		} else if matrix[row][col] > target {
			row++
		} else {
			return true
		}
	}
	return false
}

func main() {
	/**
	[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]
	5
	*/
	findNumberIn2DArray([][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}, 5)
}
