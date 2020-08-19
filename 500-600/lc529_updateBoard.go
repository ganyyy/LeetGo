package main

var Dir = [8][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func updateBoard(board [][]byte, click []int) [][]byte {
	// 从点击的地方开始, 依次DFS

	// 先看点到的地方是不是雷, 是的话直接返回
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}
	// 再看周围有几个雷
	rn, cn := len(board), len(board[0])
	var cnt int
	for _, v := range Dir {
		dx, dy := v[0]+x, v[1]+y
		if dx >= 0 && dx < rn && dy >= 0 && dy < cn {
			if board[dx][dy] == 'M' {
				cnt++
			}
		}
	}

	if cnt != 0 {
		// 填充当前位置雷的个数
		board[x][y] = byte(cnt) + '0'
	} else {
		board[x][y] = 'B'
		for _, v := range Dir {
			dx, dy := v[0]+x, v[1]+y
			if dx >= 0 && dx < rn && dy >= 0 && dy < cn {
				if board[dx][dy] == 'E' {
					updateBoard(board, []int{dx, dy})
				}
			}
		}
	}
	// 再找其他的
	return board
}
