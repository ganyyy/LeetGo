package main

import (
	"container/heap"
	"fmt"
)

type Points [][]int

func (p *Points) Len() int {
	return len(*p)
}

func (p *Points) Less(i, j int) bool {
	return less((*p)[i], (*p)[j])
}

func (p *Points) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *Points) Push(x interface{}) {
	*p = append(*p, x.([]int))
}

func (p *Points) Pop() (x interface{}) {
	x, *p = (*p)[p.Len()-1], (*p)[:p.Len()-1]
	return
}

func less(a, b []int) bool {
	return a[0]*a[0]+a[1]*a[1] < b[0]*b[0]+b[1]*b[1]
}

func kClosest(points [][]int, K int) [][]int {
	// topK 问题. 直接上堆解决 构建一个大顶堆. 元素个数为
	if len(points) <= K {
		return points
	}

	var head = points[:K]
	var p = (*Points)(&head)
	heap.Init(p)

	var t []int
	for _, v := range points[K:] {
		t = heap.Pop(p).([]int)

		if less(v, t) {
			t = v
		}
		heap.Push(p, t)
	}

	return head
}

func kClosestQuick(points [][]int, K int) [][]int {
	if len(points) <= K {
		return points
	}
	quickSelect(0, len(points)-1, points, K)
	return points[:K]
}

func dist(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

func quickSelect(l, r int, points [][]int, K int) {
	if l <= r {
		return
	}
	// 给points[right] 找到合适的位置, 保证 [l, p) 均小于 p, (p, r]均大于 p
	var p = l
	for i := l; i < r; i++ {
		if dist(points[l]) <= dist(points[r]) {
			points[l], points[p] = points[p], points[l]
			p++
		}
	}

	// 互换 r 和 p的值
	points[r], points[p] = points[p], points[r]

	// 判断是否需要继续查找
	if c := p + 1; c == K {
		// 前K个已经成功排序, 不需要进一步处理了
		return
	} else if c > K {
		// K 在左边
		quickSelect(l, c-1, points, K)
	} else {
		// K 在右边
		quickSelect(c+1, r, points, K)
	}
}

func main() {
	var src = [][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}

	fmt.Println(kClosest(src, 3))
}
