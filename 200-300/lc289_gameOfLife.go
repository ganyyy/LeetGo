package main

// 8个方向的增量
var DX = [8]int{0, 0, 1, -1, 1, 1, -1, -1}
var DY = [8]int{1, -1, 0, 0, 1, -1, 1, -1}

func gameOfLife(board [][]int) {
	// mark
	// 定义几个方向
	ln := len(board)
	if ln == 0 {
		return
	}
	lm := len(board[0])
	if lm == 0 {
		return
	}
	// 每一个int的最后一位表示当前状态
	// 每一个int的倒数第二位表示下一状态
	for i := 0; i < ln; i++ {
		for j := 0; j < lm; j++ {
			// 八个方向挨个看一下周围有几个活着的
			var cnt int
			for k := 0; k < 8; k++ {
				x, y := i+DX[k], j+DY[k]
				if uint(x) >= uint(ln) || uint(y) >= uint(lm) {
					continue
				}
				// 最后一位只能是0或者1
				cnt += board[x][y] & 1
			}
			// 分情况讨论
			// 1. 当前位置细胞是活着的
			if board[i][j]&1 == 1 {
				if cnt == 2 || cnt == 3 {
					// 下一状态是活着的
					board[i][j] = 0b11
				}
			} else if cnt == 3 {
				// 死细胞周围有三个活细胞, 会活过来
				board[i][j] = 0b10
			}
		}
	}
	// 位运算
	// 更新状态
	for i := 0; i < ln; i++ {
		for j := 0; j < lm; j++ {
			// 右移一位, 更新状态
			board[i][j] >>= 1
		}
	}
}

func main() {

}
