package main

func uniquePathsIII(grid [][]int) int {
	var cnt0, sx, sy int
	for i, row := range grid {
		for j, x := range row {
			if x == 0 {
				cnt0++
			} else if x == 1 { // 起点
				sx, sy = i, j
			}
		}
	}

	var dfs func(int, int, int) int
	// 要行进的点位, 以及剩余的点位数
	dfs = func(x, y, left int) int {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[x]) || grid[x][y] < 0 {
			return 0 // 不合法
		}
		if grid[x][y] == 2 { // 到达终点
			if left == 0 { // 必须访问所有的无障碍方格
				return 1 // 找到了一条合法路径
			}
			return 0 // 不合法
		}
		grid[x][y] = -1                   // 标记成访问过，因为题目要求「不能重复通过同一个方格」
		defer func() { grid[x][y] = 0 }() // 恢复现场
		return dfs(x-1, y, left-1) + dfs(x, y-1, left-1) +
			dfs(x+1, y, left-1) + dfs(x, y+1, left-1)
	}
	return dfs(sx, sy, cnt0+1) // +1 把起点也算上
}

// 带有记忆化的dfs
// 记忆化的关键点在于, 一个点位的访问状态是由其前面的点位决定的,
// 到达当前点位时, 如果前边访问的点位状态相同, 那么得到的结果也是相同的, 可以直接返回
func uniquePathsIII2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// 长宽加一块, 因为最多20个格子, 所以可以使用一个64位整数表示
	var vis, sx, sy int
	for i, row := range grid {
		for j, x := range row {
			if x < 0 { // 把障碍方格算上,
				vis |= 1 << (i*n + j)
			} else if x == 1 { // 起点
				sx, sy = i, j
			}
		}
	}

	all := 1<<(m*n) - 1
	type data struct{ x, y, vis int }
	memo := map[data]int{}
	var dfs func(int, int, int) int
	dfs = func(x, y, vis int) int {
		p := x*n + y
		if x < 0 || x >= m || y < 0 || y >= n || vis>>p&1 > 0 {
			return 0 // 不合法
		}
		vis |= 1 << p        // 标记访问过 (x,y)，因为题目要求「不能重复通过同一个方格」
		if grid[x][y] == 2 { // 到达终点
			if vis == all { // 必须访问所有的无障碍方格
				return 1 // 找到了一条合法路径
			}
			return 0 // 不合法
		}
		// 记忆化包含了三个元素, 这个点位坐标, 以及经过的点位
		d := data{x, y, vis}
		if v, ok := memo[d]; ok { // 之前算过
			return v
		}
		ans := dfs(x-1, y, vis) + dfs(x, y-1, vis) +
			dfs(x+1, y, vis) + dfs(x, y+1, vis)
		memo[d] = ans // 记忆化
		return ans
	}
	return dfs(sx, sy, vis)
}
