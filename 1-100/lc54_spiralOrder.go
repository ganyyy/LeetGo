package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	ln := len(matrix)
	if ln == 0 {
		return []int{}
	}
	if ln == 1 {
		return matrix[0]
	}
	ln2 := len(matrix[0])
	res := make([]int, ln*ln2)
	// 边界
	rowStart, rowEnd, colStart, colEnd := 0, ln-1, 0, ln2-1
	// 增值, 从第一行开始
	row, rowAdd, col, colAdd := 0, 0, 0, 1

	for i := 0; i < len(res); i++ {
		res[i] = matrix[row][col]
		row += rowAdd
		col += colAdd
		// 行到头了
		if col > colEnd {
			// 行+1
			rowStart += 1
			// 换行, 上到下
			col = colEnd
			colAdd = 0
			rowAdd = 1
			row += 1
			continue
		}
		if col < colStart {
			// 行-1
			rowEnd -= 1
			// 换行, 下到上
			col = colStart
			colAdd = 0
			rowAdd = -1
			row -= 1
			continue
		}
		// 列到头了
		if row > rowEnd {
			// 列-1
			colEnd -= 1
			// 换列, 右到左
			row = rowEnd
			rowAdd = 0
			colAdd = -1
			col -= 1
			continue
		}
		if row < rowStart {
			// 列+1
			colStart += 1
			// 换列, 左到右
			row = rowStart
			rowAdd = 0
			colAdd = 1
			col += 1
			continue
		}
	}
	return res
}

func main() {
	//param := [][]int {
	//	{1, 2, 3, 4},
	//	{5, 6, 7, 8},
	//	{9,10,11,12},
	//	{13,14,15,16},
	//	{17,18,19,20},
	//}
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}}))
}
