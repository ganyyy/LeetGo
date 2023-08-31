package main

const (
	// 不引入math包

	LIMIT = 0x3f3f3f3f
)

func minTrioDegree(n int, edges [][]int) int {
	// 转成位图会更省点内存诶
	g := make([][]bool, n)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]bool, n)
	}
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		g[x][y], g[y][x] = true, true
		degree[x]++
		degree[y]++
	}
	ans := LIMIT
	// 这个是优化的重点! 从小到大枚举, 保证了不会重复
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if !g[i][j] {
				continue
			}
			for k := j + 1; k < n; k++ {
				if g[i][k] && g[j][k] {
					ans = min(ans, degree[i]+degree[j]+degree[k]-6)
				}
			}
		}
	}
	if ans == LIMIT {
		return -1
	}
	return ans
}
