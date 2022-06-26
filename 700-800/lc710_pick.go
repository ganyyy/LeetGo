package main

import "math/rand"

type Solution struct {
	b2w   map[int]int
	bound int
}

func Constructor(n int, blacklist []int) Solution {
	m := len(blacklist)
	// 不排序/二分, 因为这样反而会更慢? what?
	bound := n - m
	black := map[int]bool{}
	// 难得是思路啊
	for _, b := range blacklist {
		if b >= bound {
			black[b] = true
		}
	}

	b2w := make(map[int]int, m-len(black))
	w := bound
	for _, b := range blacklist {
		if b < bound {
			for black[w] {
				w++
			}
			b2w[b] = w
			w++
		}
	}
	return Solution{b2w, bound}
}

func (s *Solution) Pick() int {
	x := rand.Intn(s.bound)
	if s.b2w[x] > 0 {
		return s.b2w[x]
	}
	return x
}
