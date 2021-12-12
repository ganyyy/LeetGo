package main

import (
	"math"
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
	// 有时候, API还不如自己写的好用

	var left, right = 0, len(tv.times)

	for left < right {
		var mid = left + (right-left)/2
		if tv.times[mid] > t {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return tv.max[left-1]
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
