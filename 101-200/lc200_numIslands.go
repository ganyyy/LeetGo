package main

func numIslands(grid [][]byte) int {
	ln := len(grid)
	if ln == 0 {
		return 0
	}
	lm := len(grid[0])
	if lm == 0 {
		return 0
	}
	var res int
	for x := 0; x < ln; x++ {
		for y := 0; y < lm; y++ {
			// 只看为1的, 一旦为1, 就会将周围为1的感染成2, 避免重复计算
			if grid[x][y] == '1' {
				infect(grid, x, y)
				res++
			}
		}
	}
	return res
}

func infect(grid [][]byte, x, y int) {
	// 越界或者不为1, 说明已经被感染或者是海洋, 直接返回即可
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] != '1' {
		return
	}
	grid[x][y] = '2'
	infect(grid, x+1, y)
	infect(grid, x-1, y)
	infect(grid, x, y+1)
	infect(grid, x, y-1)
}

func main() {

}
