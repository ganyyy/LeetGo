package main

func solve(board [][]byte) {
	ln := len(board)
	if ln == 0 {
		return
	}
	lm := len(board[0])
	if lm == 0 {
		return
	}
	// 从边界开始找 'O' 然后进行扩散
	lastR, lastC := ln-1, lm-1
	for i := 0; i < lm; i++ {
		if board[0][i] == 'O' {
			check(0, i, board)
		}
		if board[lastR][i] == 'O' {
			check(lastR, i, board)
		}
	}
	for i := 0; i < ln; i++ {
		if board[i][0] == 'O' {
			check(i, 0, board)
		}
		if board[i][lastC] == 'O' {
			check(i, lastC, board)
		}
	}
	// 所有不是1的标记为 X
	for i := 0; i < ln; i++ {
		for j := 0; j < lm; j++ {
			if board[i][j] != 1 {
				board[i][j] = 'X'
			} else {
				board[i][j] = 'O'
			}
		}
	}
}

func check(r, c int, board [][]byte) {
	ln, lm := len(board), len(board[0])
	board[r][c] = 1
	v := r - 1
	if v >= 0 && board[v][c] == 'O' {
		check(v, c, board)
	}
	v = r + 1
	if v < ln && board[v][c] == 'O' {
		check(v, c, board)
	}
	v = c - 1
	if v >= 0 && board[r][v] == 'O' {
		check(r, v, board)
	}
	v = c + 1
	if v < lm && board[r][v] == 'O' {
		check(r, v, board)
	}
}

func solve2(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	var n, m int
	n, m = len(board), len(board[0])

	// 从边界开始找 'O' 然后进行扩散
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
			return
		}
		board[x][y] = 'A'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}

	for i := 0; i < n; i++ {
		dfs(i, 0)
		dfs(i, m-1)
	}
	for i := 1; i < m-1; i++ {
		dfs(0, i)
		dfs(n-1, i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
