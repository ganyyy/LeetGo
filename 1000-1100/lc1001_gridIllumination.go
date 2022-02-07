package main

const (
	SHIFT = 16
	MASK  = (1 << SHIFT) - 1
)

func pack(x, y int) int {
	return (x << SHIFT) | y
}

func unPack(p int) (x, y int) {
	x, y = p>>SHIFT, p&MASK
	return
}

var dir = [9][2]int{
	{0, 0},
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
}

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {

	var getInc = func(p []int) (int, int) {
		var x, y = p[0], p[1]
		return (n - 1) + y - x, x + y
	}
	var ln = len(lamps)
	var row, col, inc, ainc = make(map[int]int, ln), make(map[int]int, ln), make(map[int]int, ln), make(map[int]int, ln)
	var p = make(map[int]bool, len(lamps)>>2)
	for _, lamp := range lamps {
		var x, y = lamp[0], lamp[1]
		var v = pack(x, y)
		if p[v] {
			continue
		}
		// 每个点只有一次有效(?)
		p[v] = true
		row[x]++
		col[y]++
		var i1, i2 = getInc(lamp)
		inc[i1]++
		ainc[i2]++
	}
	// fmt.Println(row, col, inc, p)
	var ret = make([]int, len(queries))
	for i, q := range queries {
		var x, y = q[0], q[1]

		var i1, i2 = getInc(q)
		if row[x] > 0 || col[y] > 0 || inc[i1] > 0 || ainc[i2] > 0 {
			ret[i] = 1
		}

		// 清空范围内的灯
		for _, d := range dir {
			var nx, ny = x + d[0], y + d[1]

			if nx < 0 || nx >= n {
				continue
			}
			if ny < 0 || ny >= n {
				continue
			}
			var v = pack(nx, ny)
			if !p[v] {
				continue
			}
			p[v] = false
			row[nx]--
			col[ny]--
			i1, i2 = getInc([]int{nx, ny})
			inc[i1]--
			ainc[i2]--
		}
	}
	return ret
}
