package main

import "slices"

func maxIncreasingCells(mat [][]int) int {
	type pair struct{ x, y int }
	g := map[int][]pair{}
	for i, row := range mat {
		for j, x := range row {
			g[x] = append(g[x], pair{i, j}) // 相同元素放在同一组，统计位置
		}
	}
	keys := make([]int, 0, len(g))
	for k := range g {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	rowMax := make([]int, len(mat))
	colMax := make([]int, len(mat[0]))
	for _, x := range keys {
		pos := g[x]
		// 值为x的所有数(不同坐标)对应的最大值
		// 因为从小往大的运算, 可以保证之后的值一定更大, 进而可以在原来的基础上直接加1
		mx := make([]int, len(pos))
		for i, p := range pos {
			mx[i] = max(rowMax[p.x], colMax[p.y]) + 1 // 先把最大值算出来，再更新 rowMax 和 colMax
		}
		for i, p := range pos {
			rowMax[p.x] = max(rowMax[p.x], mx[i]) // 更新第 p.x 行的最大 f 值
			colMax[p.y] = max(colMax[p.y], mx[i]) // 更新第 p.y 列的最大 f 值
		}
	}
	return slices.Max(colMax)
}
