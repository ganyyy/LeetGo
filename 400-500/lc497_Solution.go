package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	rects [][]int
	sum   []int
}

func Constructor(rects [][]int) Solution {
	// 每个矩形包含点的点的数量, 就是对应的权重信息
	sum := make([]int, len(rects)+1)
	for i, r := range rects {
		a, b, x, y := r[0], r[1], r[2], r[3]
		sum[i+1] = sum[i] + (x-a+1)*(y-b+1)
	}
	return Solution{rects, sum}
}

func (s *Solution) Pick() []int {
	// 随机k
	k := rand.Intn(s.sum[len(s.sum)-1])
	// 找到k对应的矩形位置
	rectIndex := sort.SearchInts(s.sum, k+1) - 1
	r := s.rects[rectIndex]

	// 再计算矩形中的点
	// y-b表示的就是矩形的高
	a, b, y := r[0], r[1], r[3]
	da := (k - s.sum[rectIndex]) / (y - b + 1)
	db := (k - s.sum[rectIndex]) % (y - b + 1)
	return []int{a + da, b + db}
}
