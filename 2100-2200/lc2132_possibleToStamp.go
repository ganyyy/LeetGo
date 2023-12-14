package main

func possibleToStampFail(grid [][]int, stampHeight int, stampWidth int) bool {
	if stampHeight <= 1 && stampWidth <= 1 {
		return true
	}
	if len(grid) == 0 || len(grid[0]) == 0 {
		return true
	}

	// 解决不了的问题:
	// 仅仅只看了行和列的连续性, 没有考虑到行和列的交叉

	const (
		ROW = 1 << (iota + 1)
		COL
	)

	defer func() {
		for _, row := range grid {
			_ = row
			// fmt.Println(row)
		}
	}()

	for _, row := range grid {
		var left int
		for right, val := range row {
			if val == 1 {
				if left != right && right-left < stampWidth {
					// 不满足连续的 stampWidth
					return false
				}
				for i := left; i < right; i++ {
					row[i] |= ROW
				}
				left = right + 1
			}
		}
		right := len(row)
		if right != left && right-left < stampWidth {
			return false
		}
		for i := left; i < right; i++ {
			row[i] |= ROW
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		var left int
		for right := 0; right < len(grid); right++ {
			if grid[right][i] == 1 {
				if left != right && right-left < stampHeight {
					return false
				}
				for j := left; j < right; j++ {
					grid[j][i] |= COL
				}
				left = right + 1
			}
		}
		right := len(grid)
		if right != left && right-left < stampHeight {
			return false
		}
		for j := left; j < right; j++ {
			grid[j][i] |= COL
		}
	}

	return true
}

func possibleToStamp(grid [][]int, stampHeight, stampWidth int) bool {
	m, n := len(grid), len(grid[0])

	// 1. 计算 grid 的二维前缀和
	s := make([][]int, m+1)
	s[0] = make([]int, n+1)
	for i, row := range grid {
		s[i+1] = make([]int, n+1)
		for j, v := range row {
			s[i+1][j+1] = s[i+1][j] + s[i][j+1] - s[i][j] + v
		}
	}

	// 2. 计算二维差分
	// 为方便第 3 步的计算，在 d 数组的最上面和最左边各加了一行（列），所以下标要 +1
	d := make([][]int, m+2)
	for i := range d {
		d[i] = make([]int, n+2)
	}
	for i2 := stampHeight; i2 <= m; i2++ {
		for j2 := stampWidth; j2 <= n; j2++ {
			i1 := i2 - stampHeight + 1
			j1 := j2 - stampWidth + 1
			if s[i2][j2]-s[i2][j1-1]-s[i1-1][j2]+s[i1-1][j1-1] == 0 {
				// 区间和为0, 说明这个区间内的值都是0, 所以这个区间整体可以放一枚邮票
				d[i1][j1]++
				d[i1][j2+1]--
				d[i2+1][j1]--
				d[i2+1][j2+1]++
			}
		}
	}

	// 3. 还原二维差分矩阵对应的计数矩阵（原地计算）
	for i, row := range grid {
		for j, v := range row {
			// 恢复区间和
			d[i+1][j+1] += d[i+1][j] + d[i][j+1] - d[i][j]
			if v == 0 && d[i+1][j+1] == 0 {
				// 如果格子上的值为0, 并且覆盖到的邮票数量为0, 说明这个格子盖不到
				return false
			}
		}
	}
	return true
}
