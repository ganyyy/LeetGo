package main

import "fmt"

func setZeroes(matrix [][]int) {
	// 判断首行和首列是否存在0
	rowFlag, colFlag := false, false
	row, col := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			colFlag = true
			break
		}
	}
	for i := 0; i < col; i++ {
		if matrix[0][i] == 0 {
			rowFlag = true
			break
		}
	}
	// 从1,1开始, 每找到一个0就把对应行首和列首置为0
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	// 遍历首行和首列, 将0对应的行和列全部置为0
	for i := 1; i < col; i++ {
		if matrix[0][i] == 0 {
			for j := 1; j < row; j++ {
				matrix[j][i] = 0
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

	// 如果存在首行或者首列的清空标志, 就将其置为0
	if rowFlag {
		for i := 0; i < col; i++ {
			matrix[0][i] = 0
		}
	}
	if colFlag {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 0, 3},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}
