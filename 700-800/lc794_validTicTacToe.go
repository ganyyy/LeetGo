package main

func validTicTacToe(board []string) bool {

	var xc, oc int
	var row, col [3]int
	var inc [2]int

	for i, r := range board {
		for j := range r {
			var c = r[j]
			var add int
			if c == 'X' {
				xc++
				add = 1
			} else if c == 'O' {
				oc++
				add = -1
			}
			row[i] += add
			col[j] += add
			if i == j {
				// 正斜线
				inc[0] += add
			}
			if i+j == 2 {
				// 反斜线
				inc[1] += add
			}
		}
	}

	if xc < oc || xc-oc > 1 {
		return false
	}

	var check = func(arr []int) bool {
		// 如果X赢了, X的数量不能和O的数量相等
		// 如果O赢了, O的数量需要和X的数量相等
		for _, v := range arr {
			if v == 3 && xc == oc {
				return false
			} else if v == -3 && xc != oc {
				return false
			}
		}
		return true
	}

	if !check(row[:]) {
		return false
	}
	if !check(col[:]) {
		return false
	}
	if !check(inc[:]) {
		return false
	}

	return true
}
