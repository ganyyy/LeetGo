//go:build ignore
// +build ignore

package main

type Node [2]int // 0: 当前值, 1: 当前最小值

type MinStack struct {
	stack []Node
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(x int) {
	if len(s.stack) == 0 {
		s.stack = append(s.stack, Node([2]int{x, x}))
		return
	}
	var last = s.stack[len(s.stack)-1]
	s.stack = append(s.stack, Node([2]int{x, min(last[1], x)}))
}

func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		return -1
	}
	var last = s.stack[len(s.stack)-1]
	return last[0]
}

func (s *MinStack) Min() int {
	if len(s.stack) == 0 {
		return -1
	}
	var last = s.stack[len(s.stack)-1]
	return last[1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
