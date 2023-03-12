package main

import "math/bits"

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	distance := make([][]uint, n)
	for _, edge := range edges {
		x, y := uint(edge[0]-1), uint(edge[1]-1)
		distance[x] = append(distance[x], y)
		x, y = y, x
		distance[x] = append(distance[x], y)
	}

	var mask, diameter uint

	var dfs func(uint) uint

	dfs = func(root uint) uint {
		var first, second uint
		mask &^= 1 << root // 去掉起点

		for _, next := range distance[root] {
			if mask&(1<<next) == 0 {
				continue
			}
			mask &^= 1 << next

			dis := 1 + dfs(next)
			if dis > first {
				first, second = dis, first
			} else if dis > second {
				second = dis
			}
		}
		if diameter < first+second {
			diameter = first + second
		}
		return first
	}

	var ret = make([]int, n-1)
	for i := uint(1); i < (1 << n); i++ {
		// 所有可能的组合
		root := uint(bits.Len(i) - 1) // 选取一个起点
		mask = i
		diameter = 0 // 这个组合中, 点位产生的距离
		dfs(root)
		if mask == 0 && diameter > 0 {
			ret[diameter-1]++
		}
	}
	return ret
}
