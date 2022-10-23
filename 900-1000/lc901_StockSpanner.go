//go:build ignore

package main

import "math"

type StockSpanner struct {
	stack [][2]int // 索引和价格
	idx   int
}

func Constructor() StockSpanner {
	return StockSpanner{[][2]int{{-1, math.MaxInt32}}, -1}
}

func (s *StockSpanner) Next(price int) int {
	s.idx++
	// 单调递减栈
	// 移除的都是连续的, 且小于当前价格的数据
	// 当前位置和最后栈顶之间的差值, 都小于当前值
	for price >= s.stack[len(s.stack)-1][1] {
		s.stack = s.stack[:len(s.stack)-1]
	}
	s.stack = append(s.stack, [2]int{s.idx, price})
	return s.idx - s.stack[len(s.stack)-2][0]
}
