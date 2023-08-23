package main

func countServers(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}

	row := make([]int, m)
	col := make([]int, n)

	var total int
	for i, r := range grid {
		for j, s := range r {
			if s == 0 {
				continue
			}
			row[i]++
			col[j]++
		}
	}
	for i, r := range grid {
		for j, s := range r {
			if s == 0 {
				continue
			}
			if row[i] > 1 || col[j] > 1 {
				total++
			}
		}
	}
	return total
}
