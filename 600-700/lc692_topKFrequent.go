package main

import (
	"fmt"
	"sort"
)

type node struct {
	v   string
	cnt int
}

type nodes []node

func (n nodes) Len() int {
	return len(n)
}

func (n nodes) Less(i, j int) bool {
	var a, b = n[i], n[j]
	if a.cnt != b.cnt {
		return a.cnt > b.cnt
	}
	return a.v < b.v
}

func (n nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func topKFrequent(words []string, k int) []string {
	var m = make(map[string]int, len(words))

	for _, v := range words {
		m[v]++
	}

	if len(m) <= k {
		var ret = make([]string, len(m))
		var n nodes = make([]node, 0, len(m))
		for key, val := range m {
			n = append(n, node{v: key, cnt: val})
		}
		sort.Sort(n)
		for i := range n {
			ret[i] = n[i].v
		}
		return ret
	}

	var n nodes = make([]node, 0, len(m))
	for key, val := range m {
		n = append(n, node{key, val})
	}

	//var n nodes = []node{
	//	{"leetcode", 1}, {"coding", 1}, {"i", 2}, {"love", 2},
	//}
	// TopK快排

	var quickSort func(st, ed int)

	quickSort = func(st, ed int) {
		fmt.Println(n)
		if st >= ed {
			return
		}
		var l, r = st + 1, ed
		for l <= r {
			for l <= r && !n.Less(st, l) {
				l++
			}
			for l <= r && n.Less(st, r) {
				r--
			}
			if l <= r {
				n.Swap(l, r)
			}
		}
		n.Swap(l-1, st)
		quickSort(st, l-2)
		if l-2 < k {
			quickSort(l, ed)
		}
	}

	quickSort(0, len(n)-1)

	var ret = make([]string, len(n))
	for i := 0; i < k; i++ {
		ret[i] = n[i].v
	}
	return ret
}

func main() {
	fmt.Println(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2))
}
