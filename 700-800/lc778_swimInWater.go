//go:build ignore

package main

import (
	"container/heap"
	"sort"
)

type pair [2]int

var dirs = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func swimInWater(grid [][]int) int {
	// 基于值二分法搞一个
	var n = len(grid)
	return sort.Search(n*n, func(cur int) bool {
		// 先卡一下, 防止出现开头大的情况
		if grid[0][0] > cur {
			return false
		}
		var visited = make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, n)
		}
		visited[0][0] = true
		var queue = []pair{{}}

		for len(queue) != 0 {
			var front = queue[0]
			queue = queue[1:]
			var i, j = front[0], front[1]
			if i == j && i == n-1 {
				return true
			}
			for _, d := range dirs {
				var x, y = d[0] + i, d[1] + j
				for 0 <= x && x < n && 0 <= y && y < n && grid[x][y] <= cur && !visited[x][y] {
					visited[x][y] = true
					queue = append(queue, pair{x, y})
				}
			}
		}
		return false
	})
}

// 算法2
func swimInWater2(grid [][]int) int {
	// 连通图+最短路径和
	var n = len(grid)

	// 轻车熟路老三样
	var fa = make([]int, n*n)
	for i := range fa {
		fa[i] = i
	}
	var find func(i int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}
	var merge = func(from, to int) {
		fa[find(from)] = find(to)
	}
	var isSet = func(a, b int) bool {
		return find(a) == find(b)
	}

	// 整合一下, 这里的边天生就是排好序的
	var pos = make([]pair, n*n)
	for i := range grid {
		for j, v := range grid[i] {
			pos[v] = pair{i, j}
		}
	}

	for i := 0; ; i++ {
		var p = pos[i]
		for _, d := range dirs {
			var x, y = p[0] + d[0], p[1] + d[1]
			if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] <= i {
				merge(x*n+y, p[0]*n+p[1])
			}
		}
		if isSet(0, n*n-1) {
			return i
		}
	}
}

type entry struct{ i, j, val int }
type hp []entry

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(entry)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func swimInWater3(grid [][]int) (ans int) {
	n := len(grid)
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	vis[0][0] = true
	h := &hp{{0, 0, grid[0][0]}}
	for {
		e := heap.Pop(h).(entry)
		ans = max(ans, e.val)
		if e.i == n-1 && e.j == n-1 {
			return
		}
		for _, d := range dirs {
			if x, y := e.i+d[0], e.j+d[1]; 0 <= x && x < n && 0 <= y && y < n && !vis[x][y] {
				vis[x][y] = true
				heap.Push(h, entry{x, y, grid[x][y]})
			}
		}
	}
}
