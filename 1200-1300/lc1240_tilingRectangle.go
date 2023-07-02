//go:build ignore

package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 感觉这题, 不如原神...
func tilingRectangle(n int, m int) int {
	// n: 行
	// m: 列
	ans := max(n, m)
	rect := make([][]bool, n)
	for i := 0; i < n; i++ {
		rect[i] = make([]bool, m)
	}

	// 以(x, y)为起点, 宽为k的矩阵是否被覆盖了
	isAvailable := func(x, y, k int) bool {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				if rect[x+i][y+j] {
					return false
				}
			}
		}
		return true
	}

	// 以(x, y)为起点, 设置宽为k的矩阵的值为val
	fillUp := func(x, y, k int, val bool) {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				rect[x+i][y+j] = val
			}
		}
	}

	// 起点是(x, y), 当前已经使用了cnt个矩形(?)
	var dfs func(int, int, int)
	dfs = func(x, y, cnt int) {
		if cnt >= ans {
			// 这是个核心啊
			// 枝减, 因为最多也就cnt个, 不能比这多了
			return
		}
		if x >= n {
			// x始终代表着矩形的起点, 如果这个超过了上限(n), 就代表着找到了结果
			ans = cnt
			return
		}
		// 检测下一行
		if y >= m {
			// m是列的上限, 这一行到头了就看下一行(或者, x,y反过来也不是不行?)
			dfs(x+1, 0, cnt)
			return
		}
		// 如当前已经被覆盖，则直接尝试下一个位置
		if rect[x][y] {
			// 这一行还没填满, 转移到这一行的下一个位置上
			dfs(x, y+1, cnt)
			return
		}
		// 上界就是剩余行列所能到达的较小值. 正方形嘛..
		for k := min(n-x, m-y); k >= 1 && isAvailable(x, y, k); k-- {
			// 将长度为 k 的正方形区域标记覆盖
			fillUp(x, y, k, true)
			// 跳过 k 个位置开始检测
			dfs(x, y+k, cnt+1)
			fillUp(x, y, k, false)
		}
	}

	dfs(0, 0, 0)
	return ans
}
