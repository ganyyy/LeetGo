package main

import (
	"fmt"
	"sort"
)

func rectangleArea(rectangles [][]int) (ans int) {
	n := len(rectangles) * 2
	// y上下边界
	hBound := make([]int, 0, n)
	for _, r := range rectangles {
		hBound = append(hBound, r[1], r[3])
	}
	fmt.Println(hBound)
	// 排序，方便下面去重
	sort.Ints(hBound)
	fmt.Println(hBound)
	// m代表不同y分段的数量
	m := 0
	for _, b := range hBound[1:] {
		if hBound[m] != b {
			m++
			hBound[m] = b
		}
	}
	// 高线的个数
	hBound = hBound[:m+1]
	// y去重
	fmt.Println(hBound)

	type tuple struct{ x, i, d int }
	sweep := make([]tuple, 0, n)
	// x左右边界
	for i, r := range rectangles {
		sweep = append(sweep, tuple{r[0], i, 1}, tuple{r[2], i, -1})
	}
	sort.Slice(sweep, func(i, j int) bool { return sweep[i].x < sweep[j].x })
	fmt.Println(sweep)

	/*
	   简单而言是这样的:
	   1. 计算竖向的分段, 并去重
	   2. 计算横向的分段, 执行过程中去重

	   根据横向分段是左边界还是右边界, 以及该边对应的, 动态的更新 seg. 给一个纵向分段都会单独计算一次
	*/

	seg := make([]int, m)
	for i := 0; i < n; i++ {
		j := i
		// 找到不同的x分段
		// 相同的x分段会被合并
		for j+1 < n && sweep[j+1].x == sweep[i].x {
			j++
		}
		if j+1 == n {
			break
		}
		// 一次性地处理掉一批横坐标相同的左右边界
		for k := i; k <= j; k++ {
			// idx:  该线对应的原始位置
			// diff: 该线带来的差值(左+1, 右-1)
			idx, diff := sweep[k].i, sweep[k].d
			// left: 下界, right: 上界(?)
			left, right := rectangles[idx][1], rectangles[idx][3]
			fmt.Println(i, j, idx, diff, left, right)
			for x := 0; x < m; x++ {
				// 计算这条边带来的收益
				// 不管怎么说, 这条纵边一定和hBound存在交集
				if left <= hBound[x] && hBound[x+1] <= right {
					seg[x] += diff
				}
			}
			fmt.Println(seg)
		}
		cover := 0
		for k := 0; k < m; k++ {
			if seg[k] > 0 {
				cover += hBound[k+1] - hBound[k]
			}
		}
		ans += cover * (sweep[j+1].x - sweep[j].x)
		i = j
	}
	return ans % (1e9 + 7)
}
