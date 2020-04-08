package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	s := []byte(word)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == s[0] {
				if dfs(i, j, board, s) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(i, j int, board [][]byte, s []byte) bool {
	if len(s) == 0 {
		return true
	}
	// 越界 或者 已被选择 或者 不相等
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == ' ' || board[i][j] != s[0] {
		return false
	} else {
		old := board[i][j]
		// 选取的标记为 ' ', 在进行回溯后 回复
		board[i][j] = ' '

		// 四个方向挨个试一下看看行不行
		res := dfs(i-1, j, board, s[1:]) ||
			dfs(i+1, j, board, s[1:]) ||
			dfs(i, j-1, board, s[1:]) ||
			dfs(i, j+1, board, s[1:])

		board[i][j] = old
		return res
	}
}

func main() {
	bs := [][]byte{
		{},
	}

	fmt.Println(exist(bs, "a"))
}
