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

var dir = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}

func numIslands2(grid [][]byte) int {
	var rowNum = len(grid)
	if rowNum == 0 {
		return 0
	}
	var colNum = len(grid[0])
	if colNum == 0 {
		return 0
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= rowNum || j < 0 || j >= colNum {
			return
		}
		if grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		for _, next := range dir {
			dfs(i+next[0], j+next[1])
		}
		return
	}
	var ret int
	for i, row := range grid {
		for j, c := range row {
			if c == '1' {
				ret++
			}
			dfs(i, j)
		}
	}

	return ret
}

var DIR = []int32{-1, 0, 1, 0, -1}

func numIslands3(grid [][]byte) int {
	row := uint(len(grid))
	if row == 0 {
		return 0
	}
	col := uint(len(grid[0]))
	if col == 0 {
		return 0
	}

	var queue, nextQueue [][2]int32
	const (
		Water   = '0'
		Land    = '1'
		Visited = '2'
	)

	var visitLands = func(x, y int32) {
		grid[x][y] = Visited
		queue = queue[:0]
		nextQueue = nextQueue[:0]
		queue = append(queue, [2]int32{x, y})

		for len(queue) != 0 {
			for _, pos := range queue {
				// 四个方向
				for ix, ay := range DIR[1:] {
					nx, ny := uint(pos[0]+DIR[ix]), uint(pos[1]+ay)
					if nx >= row || ny >= col {
						continue
					}
					if grid[nx][ny] != Land {
						continue
					}
					grid[nx][ny] = Visited
					nextQueue = append(nextQueue, [2]int32{int32(nx), int32(ny)})
				}
			}
			queue, nextQueue = nextQueue, queue[:0]
		}
	}

	var count int
	for x, r := range grid {
		for y, c := range r {
			if c != Land {
				continue
			}
			count++
			visitLands(int32(x), int32(y))
		}
	}
	return count
}

func main() {

}
