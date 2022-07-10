package main

import "math"

func cherryPickup(grid [][]int) int {
	n := len(grid)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
		for j := range f[i] {
			f[i][j] = math.MinInt32
		}
	}
	f[0][0] = grid[0][0]
	// 将整体看成是两个人一起走, 可以获取的樱桃的最大数量
	// K表示当前的步数, 上限是 n*2-1 (正方形?)
	for k := 1; k < n*2-1; k++ {
		// 优化了DP的维度, 因为当前位置的DP状态依赖于
		/*
		   向右走, x不变. 因为k/x会变, 所以y也会变
		   向下走, x-1
		*/
		// 到当前位置的x,y满足 x+y = k
		// 将整体分为上下两个半区, x1走上半区, x2走下半区
		// x1 ∈ [0, n-1]
		// x2 ∈ [x1, n-1]
		// 当x1等于n-1时, 就意味着走到了最底层
		for x1 := min(k, n-1); x1 >= max(k-n+1, 0); x1-- {
			for x2 := min(k, n-1); x2 >= x1; x2-- {
				y1, y2 := k-x1, k-x2
				if grid[x1][y1] == -1 || grid[x2][y2] == -1 {
					f[x1][x2] = math.MinInt32
					continue
				}
				res := f[x1][x2] //
				if x1 > 0 {
					res = max(res, f[x1-1][x2]) // 往下，往右
				}
				if x2 > 0 {
					res = max(res, f[x1][x2-1]) // 往右，往下
				}
				if x1 > 0 && x2 > 0 {
					res = max(res, f[x1-1][x2-1]) // 都往下
				}
				res += grid[x1][y1]
				if x2 != x1 { // 避免重复摘同一个樱桃
					res += grid[x2][y2]
				}
				f[x1][x2] = res
			}
		}
	}
	return max(f[n-1][n-1], 0)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
