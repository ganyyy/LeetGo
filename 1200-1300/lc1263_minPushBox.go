package main

func minPushBox(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var sx, sy, bx, by int // 玩家、箱子的初始位置
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == 'S' {
				sx, sy = x, y
			} else if grid[x][y] == 'B' {
				bx, by = x, y
			}
		}
	}

	ok := func(x, y int) bool { // 不越界且不在墙上
		return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] != '#'
	}
	// 做了一个转换: {0,-1}, {-1, 0}, {0, 1}, {1, 0}
	d := []int{0, -1, 0, 1, 0}

	// dp[i][j]: 玩家在位置 i, 箱子在位置 j 的最小推动次数
	dp := make([][]int, m*n)
	for i := 0; i < m*n; i++ {
		dp[i] = make([]int, m*n)
		for j := 0; j < m*n; j++ {
			dp[i][j] = 0x3f3f3f3f
		}
	}
	dp[sx*n+sy][bx*n+by] = 0              // 初始状态的推动次数为 0
	q := [][2]int{{sx*n + sy, bx*n + by}} // x*n+y
	for len(q) > 0 {
		var q1 [][2]int
		for len(q) > 0 {
			// [0]: 人所处的位置
			// [1]: 箱子所处的位置
			s1, b1 := q[0][0], q[0][1]
			q = q[1:]
			sx1, sy1, bx1, by1 := s1/n, s1%n, b1/n, b1%n
			if grid[bx1][by1] == 'T' { // 箱子已被推到目标处
				return dp[s1][b1]
			}
			for i := 0; i < 4; i++ { // 玩家向四个方向移动到另一个状态
				sx2, sy2 := sx1+d[i], sy1+d[i+1]
				s2 := sx2*n + sy2
				if !ok(sx2, sy2) { // 玩家位置不合法
					continue
				}
				if bx1 == sx2 && by1 == sy2 { // 推动箱子
					// 如果玩家下一个移动的位置和箱子的位置相同, 就尝试推箱子, 注意: 这个放入到了 p1 队列
					bx2, by2 := bx1+d[i], by1+d[i+1]
					b2 := bx2*n + by2
					// 按照BFS的思想, 这里的状态已访问是指: 之前已经有一个状态, 使得箱子在这个位置, 并且推动次数更少
					if !ok(bx2, by2) || dp[s2][b2] <= dp[s1][b1]+1 { // 箱子位置不合法 或 状态已访问
						continue
					}
					dp[s2][b2] = dp[s1][b1] + 1
					q1 = append(q1, [2]int{s2, b2})
				} else {
					// 如果玩家下一个移动的位置和箱子的位置不同, 就尝试移动玩家, 注意: 这个放入到了 p 队列
					if dp[s2][b1] <= dp[s1][b1] { // 状态已访问
						continue
					}
					// 玩家自己的移动, 不会改变箱子的位置, 所以继承了之前的状态
					dp[s2][b1] = dp[s1][b1]
					q = append(q, [2]int{s2, b1})
				}
			}
		}
		q = q1
	}
	return -1
}
