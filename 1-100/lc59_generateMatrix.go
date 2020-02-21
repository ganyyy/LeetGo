package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	// 旋转赋值
	row, rowAdd, col, colAdd := 0, 0, 0, 1
	for i := 1; i <= n*n; i++ {
		res[row][col] = i
		row += rowAdd
		col += colAdd

		// 第一行到头了
		if col >= n {
			// 向下走
			colAdd = 0
			rowAdd = 1
			row += 1
			col -= 1
			continue
		}
		// 最后一列到头了
		if row >= n {
			// 向左走
			colAdd = -1
			rowAdd = 0
			col -= 1
			row -= 1
			continue
		}
		// 最后一行到头了
		if col < 0 {
			rowAdd = -1
			colAdd = 0
			row -= 1
			col += 1
			continue
		}

		if row < n && col < n && res[row][col] != 0 {
			row -= rowAdd
			col -= colAdd
			colAdd, rowAdd = -rowAdd, colAdd
			row += rowAdd
			col += colAdd
		}
	}
	return res
}

func main() {

	res := generateMatrix(0)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}
