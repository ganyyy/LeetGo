package main

// 8向
var dir = [][2]int{
	{-1, -1},
	{-1, 0},
	{0, -1},
	{-1, 1},
	{1, -1},
	{0, 1},
	{1, 0},
	{1, 1},
}

func flipChess(chessboardSrc []string) int {
	// 获取棋盘的长和宽
	ln := len(chessboardSrc)
	lc := len(chessboardSrc[0])

	chessboard := make([][]byte, ln)
	buffChessboard := make([][]byte, ln)
	for i, row := range chessboardSrc {
		chessboard[i] = []byte(row)
		buffChessboard[i] = []byte(row)
	}

	const (
		Black = 'X'
		White = 'O'
		Space = '.'
	)

	pack := func(r, c int) [2]int { return [2]int{r, c} }

	isEnd := func(r, c int) bool { return r < 0 || r >= ln || c < 0 || c >= lc }

	var buf = make([][2]int, ln)
	buf = buf[:0]

	checkNext := func(p, d [2]int) (ret [][2]int, valid bool) {
		r, c := p[0], p[1]
		buf = buf[:0]
		for {
			r, c = r+d[0], c+d[1]
			if isEnd(r, c) {
				// 到达边界点, 还找不到黑棋, 直接pass
				valid = false
				break
			}
			chess := chessboard[r][c]
			if chess == Space {
				// 这个方向上不允许存在 Space, 否则就无法完成反转
				valid = false
				return
			}
			if chess == White {
				// 出现了一个白棋, 加入到等待队列中(不一定会反转)
				valid = true
				buf = append(buf, pack(r, c))
			} else {
				// Black
				// 出现了黑棋, 相当于到头了, 跳出循环
				break
			}
		}
		ret = buf
		return
	}

	var cnt int
	// 复用的缓冲区
	var queue [][2]int
	var next [][2]int

	bfs := func(r, c int) {

		var cur int
		queue = queue[:0]
		next = next[:0]

		// 起点就是当前点
		queue = append(queue, pack(r, c))
		// 可放置就 标成黑色的(还还原吗?)
		chessboard[r][c] = Black

		for len(queue) != 0 {
			for _, pos := range queue {
				for _, d := range dir {
					validNxt, ok := checkNext(pos, d)
					if !ok {
						continue
					}
					// 可以填充欸...
					for _, n := range validNxt {
						chessboard[n[0]][n[1]] = Black
					}
					cur += len(validNxt)
					next = append(next, validNxt...)
				}
			}
			queue, next = next, queue[:0]
		}
		if cur > cnt {
			cnt = cur
		}
		// 重置chessboard
		for i, row := range buffChessboard {
			copy(chessboard[i], row)
		}
	}

	for r, row := range chessboard {
		for c, chess := range row {
			// 坐标: r, c
			// 棋子: X/O/.
			if chess != Space {
				continue
			}
			bfs(r, c)
		}
	}

	return cnt
}
