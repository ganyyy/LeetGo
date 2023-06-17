//go:build ignore

package main

var dir = []int{0, 1, 0, -1, 0}

func closedIsland(grid [][]int) int {
	row, col := len(grid), len(grid[0])

	const (
		Valid = iota
		Bound
		Invalid
	)

	check := func(r, c int) int {
		if r < 0 || c < 0 || r >= row || c >= col {
			return Invalid
		}
		if r == 0 || r == row-1 || c == 0 || c == col-1 {
			return Bound
		}
		return Valid
	}

	const (
		Mark = 1 << 10
	)

	isMark := func(v int) bool {
		return v&Mark != 0
	}

	setMark := func(v *int) {
		*v |= Mark
	}

	unMark := func(v int) int {
		return v &^ Mark
	}

	var dfs func(r, c int) bool
	dfs = func(r, c int) (valid bool) {
		ret := check(r, c)
		if ret == Invalid {
			return true // 越界算是边界点吧?
		}
		p := &grid[r][c]
		v := *p
		n := unMark(v)
		if n == 1 {
			return true // 这算是个边界点了.
		}
		// 设置检查标记
		if ret == Bound && n == 0 {
			return false // 边界点的0是围不住的
		}
		if isMark(v) {
			return n == 0 // 如果是0, 说明是合法的
		}

		setMark(p)

		valid = true
		for ira, ca := range dir[1:] {
			ra := dir[ira]
			cc := dfs(ra+r, ca+c)
			valid = valid && cc
		}
		return valid
	}

	var ret int
	for r, row := range grid {
		for c, v := range row {
			if v != 0 {
				continue
			}
			if dfs(r, c) {
				// fmt.Println(r, c)
				ret++
			}
		}
	}

	// for _, row := range grid {
	//     for _, v := range row {
	//         fmt.Printf("%4d  ", v)
	//     }
	//     fmt.Println()
	// }

	return ret
}
