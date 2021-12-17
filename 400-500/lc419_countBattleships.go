package main

func countBattleships(board [][]byte) int {

	var ret int

	for i, row := range board {
		for j, v := range row {
			if v == '.' {
				continue
			}
			if j > 0 && board[i][j-1] == 'X' {
				continue
			}
			if i > 0 && board[i-1][j] == 'X' {
				continue
			}
			ret++
		}
	}

	return ret
}
