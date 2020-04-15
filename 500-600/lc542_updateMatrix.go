package main

import "fmt"

func updateMatrix(matrix [][]int) [][]int {
	ln := len(matrix)
	if ln == 0 {
		return matrix
	}
	lm := len(matrix[0])
	if lm == 0 {
		return matrix
	}

	// 队列
	stack := make([][2]int, 0)
	// 首先将所有的0入队, 然后将1置成-1
	for i := 0; i < ln; i++ {
		for j := 0; j < lm; j++ {
			if matrix[i][j] == 0 {
				stack = append(stack, [2]int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}

	dX := [4]int{-1, 1, 0, 0}
	dY := [4]int{0, 0, -1, 1}

	for len(stack) != 0 {
		x, y := stack[0][0], stack[0][1]
		stack = stack[1:]
		for i := 0; i < 4; i++ {
			nX, nY := x+dX[i], y+dY[i]
			// 如果是合法坐标且值为-1
			if nX >= 0 && nX < ln && nY >= 0 && nY < lm && matrix[nX][nY] == -1 {
				matrix[nX][nY] = matrix[x][y] + 1
				// 更新坐标并将新位置放入到队列中
				stack = append(stack, [2]int{nX, nY})
			}
		}
	}
	return matrix
}

const MAX = 1 << 31

func updateMatrix2(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])

	dist := make([][]int, m)
	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = MAX
		}
	}

	// left and top
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				dist[i][j] = 0
			} else {
				if i > 0 {
					dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
				}
				if j > 0 {
					dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
				}
			}
		}
	}

	// right and bottom
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i < m-1 {
				dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
			}
			if j < n-1 {
				dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
			}
		}
	}

	return dist
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func main() {
	fmt.Println(updateMatrix([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}))
}
