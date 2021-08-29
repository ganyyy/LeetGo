package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	w []int
}

func Constructor(w []int) Solution {
	for i := 1; i < len(w); i++ {
		w[i] += w[i-1]
	}
	return Solution{w: w}
}

func (m *Solution) PickIndex() int {
	var r = rand.Intn(m.w[len(m.w)-1]) + 1
	return sort.SearchInts(m.w, r)
}
