package main

import "fmt"

func totalNQueens(n int) int {
	res := 0
	chess := make([][]byte, n)
	for i := 0; i < n; i++ {
		chess[i] = make([]byte, n)
	}
	// n 行不能重复, 主对角线 row-col = const, 次对角线row+col = const
	col, main, subMain := make([]byte, n), make([]byte, 2*n), make([]byte, 2*n-1)
	var dfs func(i int)

	dfs = func(i int) {
		if i == n {
			res++
		}
		// 尝试每一列放一个皇后
		for j := 0; j < n; j++ {
			// 判断列, 主对角线, 次对脚线是否满足放置条件
			if col[j] == 0 && main[i-j+n] == 0 && subMain[i+j] == 0 {
				// 该列放入值
				chess[i][j], col[j], main[i-j+n], subMain[i+j] = 1, 1, 1, 1
				dfs(i + 1)
				// 复原
				chess[i][j], col[j], main[i-j+n], subMain[i+j] = 0, 0, 0, 0
			}
		}
	}
	dfs(0)
	return res
}

func main() {
	fmt.Println(totalNQueens(4))
}
