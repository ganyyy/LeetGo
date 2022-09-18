package main

import (
	"sort"
)

const (
	X1 = iota
	Y1
	X2
	Y2
)

type Pair struct {
	X int // 横坐标
	I int // 对应的位置
	D int // 左1/右-1
}

func rectangleArea(rectangles [][]int) int {

	n := len(rectangles) * 2
	var hBound []int
	for _, r := range rectangles {
		hBound = append(hBound, r[Y1], r[Y2])
	}
	// 纵向去重
	var l int
	{
		sort.Ints(hBound)
		for i := 1; i < len(hBound); i++ {
			if hBound[l] == hBound[i] {
				continue
			}
			l++
			hBound[l] = hBound[i]
		}
		hBound = hBound[:l+1]
	}
	sweep := make([]Pair, 0, n)
	// 横向排序
	{
		for i, r := range rectangles {
			sweep = append(sweep, Pair{
				X: r[X1],
				I: i,
				D: 1,
			}, Pair{
				X: r[X2],
				I: i,
				D: -1,
			})
		}
		sort.Slice(sweep, func(i, j int) bool { return sweep[i].X < sweep[j].X })
	}
	var ret int
	// 计算面积, 就是每个纵向分段被计算了多少次
	seg := make([]int, l)
	for i := 0; i < n; i++ {
		j := i
		// 找到下一个和当前横坐标不同的横向分段
		for j+1 < n && sweep[j+1].X == sweep[i].X {
			// 跳过相同的点
			j++
		}
		if j+1 == n {
			// 到达末尾了
			break
		}
		// 累加x对应的所有纵向分段计数
		for k := i; k <= j; k++ {
			idx, d := sweep[k].I, sweep[k].D

			bottom, top := rectangles[idx][Y1], rectangles[idx][Y2]

			for x := 0; x < l; x++ {
				if bottom <= hBound[x] && hBound[x+1] <= top {
					seg[x] += d
				}
			}
		}
		// 计算面积
		var cover int
		for k := 0; k < l; k++ {
			if seg[k] > 0 {
				// 只要大于0, 就表示当前纵向分段是被覆盖的
				cover += hBound[k+1] - hBound[k]
			}
		}
		// 长是cover, 宽是X'-X
		ret += cover * (sweep[j+1].X - sweep[j].X)
		i = j
	}
	return ret % (1e9 + 7)
}
