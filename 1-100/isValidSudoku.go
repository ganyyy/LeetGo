package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rM := make([][]byte, 9)
	cM := make([][]byte, 9)
	bM := make([][]byte, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			v := board[i][j]
			if v == '.' {
				continue
			}
			v = v - '1'
			r := rM[i]
			if r == nil {
				r = make([]byte, 9)
				rM[i] = r
			}
			if r[v] != 0 {
				return false
			} else {
				r[v] = 1
			}

			c := cM[j]
			if c == nil {
				c = make([]byte, 9)
				cM[j] = c
			}
			if c[v] != 0 {
				return false
			} else {
				c[v] = 1
			}
			bb := (i/3)*3 + j/3
			b := bM[bb]
			if b == nil {
				b = make([]byte, 9)
				bM[bb] = b
			}
			if b[v] != 0 {
				return false
			} else {
				b[v] = 1
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}
