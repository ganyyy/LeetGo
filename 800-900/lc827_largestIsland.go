package main

var DIR = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func largestIsland(grid [][]int) int {
	// 回调不见得性能好, 但是看起来爽啊
	// 真的不如rust的0成本抽象

	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	if col == 0 {
		return 0
	}

	// 标记对应的大小
	var size []int

	// 每个联通的岛屿的对应标记
	flag := make([][]int, row)
	for i := range flag {
		flag[i] = make([]int, col)
	}

	valid := func(nr, nc int) bool {
		if nr < 0 || nr >= row {
			return false
		}
		if nc < 0 || nc >= col {
			return false
		}
		return true
	}

	iterator := func(cb func(r, c int)) {
		for r := 0; r < row; r++ {
			for c := 0; c < col; c++ {
				cb(r, c)
			}
		}
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dirIterator := func(r, c int, cb func(nr, nc int)) {
		for _, d := range DIR {
			nr, nc := r+d[0], c+d[1]
			if !valid(nr, nc) {
				continue
			}
			cb(nr, nc)
		}
	}

	var cnt int
	var dfs func(int, int)
	dfs = func(r, c int) {
		if flag[r][c] != 0 {
			return
		}
		flag[r][c] = len(size) + 1
		cnt++
		dirIterator(r, c, func(nr, nc int) {
			if grid[nr][nc] == 0 {
				return
			}
			dfs(nr, nc)
		})
	}

	// 只要不为grid不为空, 那么最少占用一个面积
	ret := 1

	// 记录每个岛屿的大小
	iterator(func(r, c int) {
		if grid[r][c] == 0 {
			return
		}
		if flag[r][c] != 0 {
			return
		}
		cnt = 0
		dfs(r, c)
		size = append(size, cnt)
		ret = max(ret, cnt)
	})

	// 根据填充为0的位置, 并计算最大值
	iterator(func(r, c int) {
		if grid[r][c] != 0 {
			return
		}
		s := 1
		var all [4]int // 四个方向上选取的区域
		var idx int
		dirIterator(r, c, func(nr, nc int) {
			f := flag[nr][nc]
			if f == 0 {
				return
			}
			for i := 0; i < idx; i++ {
				// 不进行重复的选取
				if all[i] == f {
					return
				}
			}
			all[idx] = f
			idx++
			s += size[f-1]
		})
		ret = max(ret, s)
	})
	return ret
}
