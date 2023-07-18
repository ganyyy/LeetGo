//go:build ignore

package main

func robotSim(commands []int, obstacles [][]int) int {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	px, py, d := 0, 0, 1
	set := make(map[int]bool)
	for _, obstacle := range obstacles {
		// [2]int 做key, 性能远差于 int*60001+int
		set[obstacle[0]*60001+obstacle[1]] = true
	}
	res := 0
	for _, c := range commands {
		if c < 0 {
			if c == -1 {
				d = (d + 1) % 4
			} else {
				d = (d + 3) % 4
			}
		} else {
			for i := 0; i < c; i++ {
				if set[(px+dirs[d][0])*60001+py+dirs[d][1]] {
					break
				}
				px += dirs[d][0]
				py += dirs[d][1]
				res = max(res, px*px+py*py)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
