package main

var dir = []int{0, 1, 0, -1, 0}

func nearestExit(maze [][]byte, entrance []int) int {
	var cur = [][2]int{{entrance[0], entrance[1]}}
	var next [][2]int
	var step int

	row, col := len(maze), len(maze[0])

	var isEnd = func(a, b int) bool {
		return a == 0 || a == row-1 || b == 0 || b == col-1
	}
	maze[entrance[0]][entrance[1]] = '|'

	for len(cur) > 0 {
		step++
		for _, p := range cur {
			for i := range dir[1:] {
				nr, nc := p[0]+dir[i], p[1]+dir[i+1]
				if nc < 0 || nc >= col || nr < 0 || nr >= row {
					continue
				}
				b := maze[nr][nc]
				if b != '.' {
					continue
				}
				maze[nr][nc] = '|'
				if isEnd(nr, nc) {
					return step
				}
				next = append(next, [2]int{nr, nc})
			}
		}
		cur, next = next, cur[:0]
	}
	return -1
}
