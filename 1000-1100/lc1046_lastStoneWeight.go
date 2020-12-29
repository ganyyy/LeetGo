package main

import (
	"container/heap"
)

type Stones []int

func (s *Stones) Len() int {
	return len(*s)
}

func (s *Stones) Less(i, j int) bool {
	return (*s)[i] > (*s)[j]
}

func (s *Stones) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *Stones) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *Stones) Pop() (x interface{}) {
	x, *s = (*s)[s.Len()-1], (*s)[:s.Len()-1]
	return
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	if len(stones) == 2 {
		var res = stones[0] - stones[1]
		if res < 0 {
			return -res
		}
	}

	var s = (*Stones)(&stones)
	heap.Init(s)

	for s.Len() >= 2 {
		var x, y = heap.Pop(s).(int), heap.Pop(s).(int)
		if x > y {
			heap.Push(s, x-y)
		}
	}

	if s.Len() == 1 {
		return (*s)[0]
	} else {
		return 0
	}
}

func main() {
	println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}
