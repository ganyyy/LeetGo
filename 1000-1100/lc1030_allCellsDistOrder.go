package main

import "sort"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	res := make([][]int, R*C)
	res[0] = []int{r0, c0}
	r, c := r0, c0
	for step := 1; step < R*C; {
		// 从上方开始, 按照顺时针方向走
		r--
		for r < r0 {
			if 0 <= r && c < C {
				res[step] = []int{r, c}
				step++
			}
			r++
			c++
		}
		for c > c0 {
			if r < R && c < C {
				res[step] = []int{r, c}
				step++
			}
			r++
			c--
		}
		for r > r0 {
			if r < R && 0 <= c {
				res[step] = []int{r, c}
				step++
			}
			r--
			c--
		}
		for c < c0 {
			if 0 <= r && 0 <= c {
				res[step] = []int{r, c}
				step++
			}
			r--
			c++
		}
	}
	return res

}

func allCellsDistOrderSort(R int, C int, r0 int, c0 int) [][]int {
	// 当前点的距离是0, 然后向四周扩散

	var res = make([][]int, R*C)

	// 最大跨度是已知的, 就是 (r0, c0) 距离 (0,0)和(R-1,C-1)之间的最大曼哈顿距离

	// 以最大距离为跨度, 分别加入合法的坐标值

	// 或者直接把所有点都插入到数组中, 在搞一个 排序...
	var cnt int
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			res[cnt] = []int{r, c}
			cnt++
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return dis(res[i][0], res[i][1], r0, c0) < dis(res[j][0], res[j][1], r0, c0)
	})

	return res
}

func dis(r1, c1, r2, c2 int) int {
	return abs(r1-r2) + abs(c1-c2)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
