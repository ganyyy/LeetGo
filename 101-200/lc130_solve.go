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
