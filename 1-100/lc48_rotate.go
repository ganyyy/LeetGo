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
func rotate2(matrix [][]int) {
	n := len(matrix)
	// 对角线反转
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 中点反转
	mid, t := n>>1, n-1
	for i := 0; i < n; i++ {
		for j := 0; j < mid; j++ {
			matrix[i][j], matrix[i][t-j] = matrix[i][t-j], matrix[i][j]
		}
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
