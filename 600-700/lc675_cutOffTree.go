package main

import "sort"

func cutOffTree(forest [][]int) (ans int) {
	var dir4 = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	type Pair struct{ dis, x, y int }
	var trees []Pair
	for i, row := range forest {
		for j, h := range row {
			if h > 1 {
				trees = append(trees, Pair{h, i, j})
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool { return trees[i].dis < trees[j].dis })

	bfs := func(sx, sy, tx, ty int) int {
		m, n := len(forest), len(forest[0])
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[sx][sy] = true
		q := []Pair{{0, sx, sy}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.x == tx && p.y == ty {
				return p.dis
			}
			for _, d := range dir4 {
				if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && forest[x][y] > 0 {
					vis[x][y] = true
					q = append(q, Pair{p.dis + 1, x, y})
				}
			}
		}
		return -1
	}

	preX, preY := 0, 0
	for _, t := range trees {
		d := bfs(preX, preY, t.x, t.y)
		if d < 0 {
			return -1
		}
		ans += d
		preX, preY = t.x, t.y
	}
	return
}
