//go:build ignore
// +build ignore

package main

import (
	. "leetgo/data"
	"math/rand"
)

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

func (s *Solution) GetRandom() int {
	// 蓄水池抽样

	var res = s.head.Val

	var idx = 2
	for cur := s.head.Next; cur != nil; cur = cur.Next {
		// 每个点被抽取的概率都是 1/n
		if rand.Intn(idx) == 0 {
			res = cur.Val
		}
		idx++
	}

	return res
}
