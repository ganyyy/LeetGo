package main

func movesToChessboard(board [][]int) int {
	length := len(board) // 如果可以组合成一个有效的棋盘,
	// 那么相邻两行/两列相互异或的结果一定等同于 (1<<n)-1

	// step1: 判断是否满足 其四个角相互异或的结果一定是0(0/1的数量相同)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if board[0][0]^board[i][0]^board[0][j]^board[i][j] == 1 {
				return -1
			}
		}
	}

	// step2: 判断行/列中1/0的个数
	// 如果是偶数, 则各占一半; 否则0/1会多一个
	// 只看第一行/列, 划重点, 要考的

	// 每次移动, 都会导致两位0/1发生变化. 综上所述, 对比有效棋盘的位差/2
	// 就是需要移动的次数
	var rowCnt, colCnt int   // 首行/首列1的个数
	var rowSwap, colSwap int // 行/列移动的计数
	for i := 0; i < length; i++ {
		rowCnt += board[0][i]
		colCnt += board[i][0]

		// 假定首个位置应该是1, 那么序列就是 1010101 ...
		// 所以当 x & 1 == i & 1 时(i&1 => 01010101),
		// 就交换一下(其实如果swap == n, 那就不用交换)
		if board[i][0]&1 == i&1 {
			rowSwap++
		}
		if board[0][i]&1 == i&1 {
			colSwap++
		}
	}

	// 判断个数
	if rowCnt != length/2 && rowCnt != (length+1)/2 {
		return -1
	}
	if colCnt != length/2 && colCnt != (length+1)/2 {
		return -1
	}

	// step3: 计算需要移动的次数
	if length&1 == 0 {
		// 偶数
		rowSwap = min(rowSwap, length-rowSwap)
		colSwap = min(colSwap, length-colSwap)
	} else {
		// 奇数转偶数
		if rowSwap&1 == 1 {
			rowSwap = length - rowSwap
		}
		if colSwap&1 == 1 {
			colSwap = length - colSwap
		}
	}

	return (rowSwap + colSwap) / 2
}
