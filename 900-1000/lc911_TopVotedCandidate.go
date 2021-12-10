//go:build ignore
// +build ignore

package main

import (
	"math"
	"sort"
)

type TopVotedCandidate struct {
	max   []int
	times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	var m = make(map[int]int)
	var max = make([]int, 0, len(times))

	var mm = math.MinInt32
	var mp int
	for i := range times {
		var p = persons[i]
		var old = m[p]
		if old+1 >= mm {
			mm = old + 1
			mp = p
		}
		m[p] = old + 1
		max = append(max, mp)
	}

	return TopVotedCandidate{
		max:   max,
		times: times,
	}
}

func (tv *TopVotedCandidate) Q(t int) int {
	var idx = sort.Search(len(tv.max), func(i int) bool {
		return t <= tv.times[i]
	})

	if idx >= len(tv.max) {
		return tv.max[idx-1]
	}
	if tv.times[idx] > t {
		if idx > 0 {
			return tv.max[idx-1]
		} else {
			return tv.max[0]
		}
	}
	return tv.max[idx]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */

/**
["TopVotedCandidate","q","q","q","q","q","q"]
[[[0,1,1,0,0,1,0],[0,5,10,15,20,25,30]],[3],[12],[25],[15],[24],[8]]
*/

func main() {
	var t = Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})

	println(t.Q(12))
	println(t.Q(10))
	println(t.Q(15))
}
