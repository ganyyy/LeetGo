//go:build ignore

package main

import "math/rand"

type Solution struct {
	Elem map[int]struct{}
	Cnt  int
	M, N int
}

func Constructor(m int, n int) Solution {
	return Solution{
		Elem: make(map[int]struct{}),
		Cnt:  m * n,
		M:    m,
		N:    n,
	}
}

func (this *Solution) Flip() []int {
	for {
		var idx = rand.Intn(this.Cnt)
		if _, ok := this.Elem[idx]; ok {
			continue
		}
		this.Elem[idx] = struct{}{}
		return []int{idx / this.N, idx % this.N}
	}
}

func (this *Solution) Reset() {
	for i := range this.Elem {
		delete(this.Elem, i)
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(m, n);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
