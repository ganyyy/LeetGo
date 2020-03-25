package main

import "fmt"

func numRookCaptures(board [][]byte) int {
	var row, col, res int
end:
	for ; row < 8; row++ {
		for col = 0; col < 8; col++ {
			if board[row][col] == 'R' {
				break end
			}
		}
	}
	// 右边
	for i := col + 1; i < 8; i++ {
		if board[row][i] == 'B' {
			break
		}
		if board[row][i] == 'p' {
			res++
			break
		}
	}
	// 左边
	for i := col - 1; i > -1; i-- {
		if board[row][i] == 'B' {
			break
		}
		if board[row][i] == 'p' {
			res++
			break
		}
	}
	// 上边
	for i := row - 1; i > -1; i-- {
		if board[i][col] == 'B' {
			break
		}
		if board[i][col] == 'p' {
			res++
			break
		}
	}
	// 下边
	for i := row + 1; i < 8; i++ {
		if board[i][col] == 'B' {
			break
		}
		if board[i][col] == 'p' {
			res++
			break
		}
	}
	return res
}

func main() {
	ret := numRookCaptures([][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', 'p', 'p', '.', '.', '.', '.'},
		{'.', 'p', 'p', 'R', '.', 'p', '.', 'p'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', 'p', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	})
	fmt.Println(ret)
}
