package main

func findBall(grid [][]int) []int {
	var c = len(grid[0])
	var r = len(grid)
	var ret = make([]int, c)
	for i := 0; i < c; i++ {
		var cc = i
		for j := 0; j < r; j++ {
			var dir = grid[j][cc]
			cc += dir
			if cc < 0 || cc >= c || grid[j][cc] != dir {
				cc = -1
				break
			}
		}
		ret[i] = cc
	}

	return ret
}
