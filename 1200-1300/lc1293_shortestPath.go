package main

func shortestPath(grid [][]int, k int) int {
	var DIR = []int{-1, 0, 1, 0, -1}

	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	if col == 0 {
		return 0
	}
	type Pos struct {
		X, Y, Remain int
	}

	if grid[0][0] == 1 {
		if k == 0 {
			// 特殊情况: 起点是障碍, 并且k=0
			return -1
		}
		k--
	}
	// 特殊情况: 只有一个点
	if row == 1 && col == 1 {
		return 0
	}

	const (
		Shift = 10               // k的偏移量
		Mask  = (1 << Shift) - 1 // 用于获取k的值
	)

	// 标记访问状态
	set := func(x, y, k int) { grid[x][y] |= (1 + k) << Shift }

	// 获取访问状态: (k, val)
	get := func(v int) (int, int) { return (v >> Shift) - 1, v & Mask }

	var curQueue, nextQueue []Pos
	curQueue = append(curQueue, Pos{X: 0, Y: 0, Remain: k})
	set(0, 0, k)
	var step int
	for len(curQueue) != 0 {
		step++
		for _, pos := range curQueue {
			for i, addX := range DIR[1:] {
				addY := DIR[i]
				newX, newY := pos.X+addX, pos.Y+addY
				if uint(newX) >= uint(row) || uint(newY) >= uint(col) {
					continue
				}
				oldK, val := get(grid[newX][newY])

				if oldK >= pos.Remain {
					// 当前位置已经访问过, 而且保留的k大于等于当前的k, 那么就不需要继续了
					continue
				}
				var newK = pos.Remain
				if val == 1 {
					// 当前点是个障碍
					if newK == 0 {
						// 但是无法继续破墙
						continue
					}
					newK--
				}
				if newX == row-1 && newY == col-1 {
					// 找到了终点, 直接返回
					return step
				}
				if oldK < newK {
					// 两种情况需要更新
					// 1: 没有迭代过(-1)
					// 2: 这个位置上保留的k小于继承而来的k, 那么就重走一遍
					set(newX, newY, newK)
					nextQueue = append(nextQueue, Pos{
						X: newX, Y: newY, Remain: newK,
					})
				}
			}
		}
		curQueue, nextQueue = nextQueue, curQueue[:0]
	}
	return -1
}
