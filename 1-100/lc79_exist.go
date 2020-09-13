package main

var dir = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

const pass = '0'

func exist(board [][]byte, word string) bool {

	if len(board) == 0 {
		return false
	}
	ln := len(board)
	if len(board[0]) == 0 {
		return false
	}
	lm := len(board[0])

	if len(word) == 0 {
		return false
	}
	lw := len(word)

	var helper func(x, y, idx int) bool
	helper = func(x, y, idx int) bool {
		if idx == lw {
			return true
		}
		if x < 0 || x >= ln {
			return false
		}
		if y < 0 || y >= lm {
			return false
		}
		if board[x][y] != word[idx] {
			return false
		}
		board[x][y] = pass
		defer func() {
			board[x][y] = word[idx]
		}()
		for _, v := range dir {
			if helper(x+v[0], y+v[1], idx+1) {
				return true
			}
		}
		return false
	}

	for r, row := range board {
		for c := range row {
			if helper(r, c, 0) {
				return true
			}
		}
	}

	return false
}

func main() {

}
