package main

import (
	"container/heap"
	"sort"
)

type Pair int

func (p Pair) Index() int { return int(p) }
func (p Pair) Value() int { return int(p) }

func NewPair(idx, val int) Pair { return Pair(val) }

type Pairs []Pair

func (p Pairs) Len() int            { return len(p) }
func (p Pairs) Less(i, j int) bool  { return p[i] < p[j] }
func (p Pairs) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *Pairs) Push(x interface{}) { *p = append(*p, x.(Pair)) }
func (p *Pairs) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[:n-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	ln := len(costs)
	if ln <= candidates*2 || ln <= k {
		sort.Ints(costs)
		var ret int
		for i := 0; i < k && i < ln; i++ {
			ret += costs[i]
		}
		return int64(ret)
	}

	pairs := make([]Pair, 0, ln)
	for i, v := range costs {
		pairs = append(pairs, NewPair(i, v))
	}

	// 前后双指针

	var cnt int
	var total int
	var li, ri = candidates, ln - candidates
	var left Pairs = pairs[:li]
	var right Pairs = pairs[ri:]
	heap.Init(&left)
	heap.Init(&right)
	for cnt < k && left.Len() > 0 && right.Len() > 0 {
		cnt++
		if left[0].Value() <= right[0].Value() {
			total += left[0].Value()
			heap.Pop(&left)
			if li < ri {
				heap.Push(&left, pairs[li])
				li++
			}
		} else {
			total += right[0].Value()
			heap.Pop(&right)
			if li < ri {
				ri--
				heap.Push(&right, pairs[ri])
			}
		}
	}
	if cnt < k {
		remain := left
		if remain.Len() == 0 {
			remain = right
		}
		sort.Sort(remain)
		for i := 0; i < (k-cnt) && i < remain.Len(); i++ {
			total += int(remain[i])
		}
	}
	return int64(total)
}
