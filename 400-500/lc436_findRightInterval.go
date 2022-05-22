package main

import "sort"

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	type pair struct{ x, i int }
	starts := make([]pair, n)
	ends := make([]pair, n)
	for i, p := range intervals {
		starts[i] = pair{p[0], i}
		ends[i] = pair{p[1], i}
	}
	// 起点和终点的正序
	sort.Slice(starts, func(i, j int) bool { return starts[i].x < starts[j].x })
	sort.Slice(ends, func(i, j int) bool { return ends[i].x < ends[j].x })

	ans := make([]int, n)
	j := 0
	for _, p := range ends {
		// 找到首个 起点 >= 终点的值, 这个就是
		for j < n && starts[j].x < p.x {
			j++
		}
		if j < n {
			ans[p.i] = starts[j].i
		} else {
			ans[p.i] = -1
		}
	}
	return ans
}

type Pair [2]int

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func findRightInterval2(intervals [][]int) []int {
	var length = len(intervals)

	var startPair, endPair PairList = make([]Pair, 0, length), make([]Pair, 0, length)

	for i, interval := range intervals {
		startPair = append(startPair, [2]int{interval[0], i})
		endPair = append(endPair, [2]int{interval[1], i})
	}

	sort.Sort(startPair)
	sort.Sort(endPair)

	// fmt.Println(startPair)
	// fmt.Println(endPair)

	var ret = make([]int, length)

	var start int
	for _, end := range endPair {
		for start < length && startPair[start][0] < end[0] {
			start++
		}
		var pos = -1
		if start < length {
			pos = startPair[start][1]
		}
		ret[end[1]] = pos
	}

	return ret
}
