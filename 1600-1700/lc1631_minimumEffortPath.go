package main

import (
	"container/heap"
	"math"
	"sort"
)

type pair [2]int

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPath(heights [][]int) int {
	// 好家伙, 还是图论.
	var l1, l2 int
	if l1 = len(heights); l1 == 0 {
		return 0
	}
	if l2 = len(heights[0]); l2 == 0 {
		return 0
	}
	// 好家伙, 学到了. sort.Search()

	// 值二分
	// 从1e6 开始执行二分查找. 查找的条件是 该值作为最大高度差时可以满足从左上角到右下角
	return sort.Search(1e6, func(maxHeightDiff int) bool {
		var vis = make([][]bool, l1)
		for i := range vis {
			vis[i] = make([]bool, l2)
		}
		// 标记一下
		vis[0][0] = true
		// 查找的队列
		var queue = []pair{{}}
		for len(queue) != 0 {
			var front = queue[0]
			queue = queue[1:]
			// 说明在当前的最大高度差情况下, 已经可以走到根节点了,
			// 尝试继续向更小的高度差进行搜索
			if front[0] == l1-1 && front[1] == l2-1 {
				return true
			}

			// 四个方向执行广度优先遍历
			for _, d := range dirs {
				var x, y = d[0] + front[0], d[1] + front[1]
				if 0 <= x && x < l1 && 0 <= y && y < l2 {
					if !vis[x][y] && abs(heights[x][y], heights[front[0]][front[1]]) <= maxHeightDiff {
						vis[x][y] = true
						queue = append(queue, pair{x, y})
					}
				}
			}
		}
		// 当前的高度差不足以走到右下角, 返回false
		return false
	})
}

type edge [3]int

// 解法2, 并查集
func minimumEffortPath2(heights [][]int) int {

	var n, m = len(heights), len(heights[0])
	// 那啥 k 算法. 求最小生成树的那个
	var fa = make([]int, n*m)
	for i := range fa {
		fa[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if i != fa[i] {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}

	var merge = func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	var isSame = func(a, b int) bool {
		return find(a) == find(b)
	}

	// 每条边的权重为 相邻两个点的高度差绝对值
	var edges = make([]edge, 0, m*n)
	for i, row := range heights {
		for j, h := range row {
			id := i*m + j
			if i > 0 {
				edges = append(edges, edge{id - m, id, abs(h, heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(h, heights[i][j-1])})
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	// 合并端点.
	for _, e := range edges {
		merge(e[0], e[1])
		// 如果此时左上和右下连接成功了, 就返回当前边的高度差
		if isSame(0, n*m-1) {
			return e[2]
		}
	}
	return 0
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

type point struct{ x, y, maxDiff int }
type hp []point

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].maxDiff < h[j].maxDiff }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(point)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

type pair2 struct{ x, y int }

var dir4 = []pair2{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 解法3, 最短路径.(和最小生成树啥关系?)
func minimumEffortPath3(heights [][]int) int {
	n, m := len(heights), len(heights[0])

	// 初始化 定义每个点的最大高度差
	maxDiff := make([][]int, n)
	for i := range maxDiff {
		maxDiff[i] = make([]int, m)
		for j := range maxDiff[i] {
			maxDiff[i][j] = math.MaxInt64
		}
	}
	maxDiff[0][0] = 0

	// 搞个堆
	h := &hp{{}}
	for {
		// 从堆里找出高度差最小的值
		p := heap.Pop(h).(point)
		// 如果找到了, 返回即可
		if p.x == n-1 && p.y == m-1 {
			return p.maxDiff
		}
		// 如果当前点的高度差大于记录的最小的高度差, 说明已经处理过了
		if maxDiff[p.x][p.y] < p.maxDiff {
			continue
		}
		for _, d := range dir4 {
			if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m {
				if diff := max(p.maxDiff, abs(heights[x][y], heights[p.x][p.y])); diff < maxDiff[x][y] {
					// 这里上个标记, 并且将最新的高度差记录下来
					maxDiff[x][y] = diff
					heap.Push(h, point{x, y, diff})
				}
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
