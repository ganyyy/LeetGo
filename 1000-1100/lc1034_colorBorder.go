package main

import "fmt"

var dir = [4][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	// 查看给定点是否处于连通分量的中间
	var m, n = len(grid), len(grid[0])
	var set = make([]bool, m*n)

	var dfs func(i, j int)

	var parseKey = func(i, j int) int {
		return i*n + j
	}

	var abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	dfs = func(i, j int) {
		var key = parseKey(i, j)
		if set[key] {
			return
		}
		set[key] = true
		var cur = abs(grid[i][j])
		var board bool
		// 存在一个不相等的颜色, 或者到达边界. 那么该块就需要上色
		for _, d := range dir {
			var ii, jj = i + d[0], j + d[1]
			if ii >= 0 && ii < m && jj >= 0 && jj < n {
				if abs(grid[ii][jj]) == cur {
					if !set[parseKey(ii, jj)] {
						dfs(ii, jj)
					}
				} else {
					board = true
				}
			} else {
				board = true
			}
		}

		if board {
			grid[i][j] = -cur
		}
	}

	dfs(row, col)

	for _, row := range grid {
		for j, c := range row {
			if c < 0 {
				row[j] = color
			}
		}
	}

	return grid
}

func main() {
	fmt.Println(colorBorder([][]int{
		{1, 2, 2},
		{2, 3, 2},
	}, 0, 1, 3))
}
