package main

import "fmt"

func rotate(matrix [][]int) {
	count := len(matrix) - 1
	if count <= 0 {
		return
	}
	for i := 0; ; i++ {
		localRotate(matrix, i, count)
		count -= 2
		if count < 0 {
			break
		}
	}
}

func localRotate(matrix [][]int, start, len int) {
	for i := 0; i < len; i++ {
		// 核心是这一块, 转置的位置关系
		matrix[i+start][start+len],
			matrix[start+len][start+len-i],
			matrix[start+len-i][start],
			matrix[start][i+start] =
			matrix[start][i+start],
			matrix[i+start][start+len],
			matrix[start+len][start+len-i],
			matrix[start+len-i][start]
	}
}

func main() {
	var matrix = [][]int{
		{1, 2},
		{4, 5},
	}
	rotate(matrix)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
}
