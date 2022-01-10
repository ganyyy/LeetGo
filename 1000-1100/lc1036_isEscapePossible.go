package main

import "sort"

type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func isEscapePossible(block [][]int, source, target []int) bool {
	const (
		blocked = -1 // 在包围圈中
		valid   = 0  // 不在包围圈中
		found   = 1  // 无论在不在包围圈中，但在 n(n-1)/2 步搜索的过程中经过了 target

		boundary int = 1e6
	)

	n := len(block)
	// 小于两个的block无法围住
	if n < 2 {
		return true
	}

	// 所有的block节点
	blockSet := map[pair]bool{}
	for _, b := range block {
		blockSet[pair{b[0], b[1]}] = true
	}

	check := func(start, finish []int) int {
		sx, sy := start[0], start[1]
		fx, fy := finish[0], finish[1]
		// 探寻的上限: 最优的阻拦方法是每一行每一列只有一个block
		// 那么其可以围住的最大的空间为一个长度为n的等腰直角三角形
		countdown := n * (n - 1) / 2

		// 可选择的路径
		q := []pair{{sx, sy}}
		// 已探寻的路径
		vis := map[pair]bool{{sx, sy}: true}
		for len(q) > 0 && countdown > 0 {
			p := q[0]
			q = q[1:]
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				np := pair{x, y}
				if 0 <= x && x < boundary && 0 <= y && y < boundary && !blockSet[np] && !vis[np] {
					if x == fx && y == fy {
						// 找到了终点. 此时可能二者都在包围圈内
						return found
					}
					countdown--
					vis[np] = true
					q = append(q, np)
				}
			}
		}

		// 这种情况说明: 没有可以行走的下一步, 说明在包围圈内
		if countdown > 0 {
			return blocked
		}
		// 围不住了
		return valid
	}

	res := check(source, target)
	// 这里为啥需要反向查询一次呢? 是因为围不住起点->终点, 还需要保证围不住终点->起点. 这样才能证明二者的连通性
	return res == found || res == valid && check(target, source) != blocked
}

// 离散化 a，返回的哈希表中的键值对分别为 a 中的原始值及其离散化后的对应值
// 以及离散化后, a可达的最远位置
func discrete(a []int) (map[int]int, int) {
	sort.Ints(a)

	id := 0
	if a[0] > 0 {
		id = 1
	}
	mapping := map[int]int{a[0]: id}
	pre := a[0]
	for _, v := range a[1:] {
		if v != pre {
			if v == pre+1 {
				// 紧挨着时, id递增
				id++
			} else {
				// 否则, 为了避免离散后二者变成相邻的点为
				// 中间需要空出一个值
				id += 2
			}
			mapping[v] = id
			pre = v
		}
	}

	// 看末尾的值是否达到了边界.
	// 如果没到达边界, 那么id需要自增一下(离散后可以达到的最远的位置+1)
	const boundary int = 1e6
	if a[len(a)-1] != boundary-1 {
		id++
	}

	return mapping, id
}

func isEscapePossible2(block [][]int, source, target []int) bool {
	n := len(block)
	if n < 2 {
		return true
	}
	rows := []int{source[0], target[0]}
	cols := []int{source[1], target[1]}
	for _, b := range block {
		rows = append(rows, b[0])
		cols = append(cols, b[1])
	}

	// 离散化行列坐标
	rMapping, rBound := discrete(rows)
	cMapping, cBound := discrete(cols)

	// 经过离散化整理后, 所有的关键点
	grid := make([][]bool, rBound+1)
	for i := range grid {
		grid[i] = make([]bool, cBound+1)
	}
	for _, b := range block {
		grid[rMapping[b[0]]][cMapping[b[1]]] = true
	}

	sx, sy := rMapping[source[0]], cMapping[source[1]]
	tx, ty := rMapping[target[0]], cMapping[target[1]]
	grid[sx][sy] = true
	// 常规的BFS
	q := []pair{{sx, sy}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		for _, d := range dirs {
			x, y := p.x+d.x, p.y+d.y
			if 0 <= x && x <= rBound && 0 <= y && y <= cBound && !grid[x][y] {
				if x == tx && y == ty {
					return true
				}
				grid[x][y] = true
				q = append(q, pair{x, y})
			}
		}
	}
	return false
}
