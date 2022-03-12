//go:build ignore
// +build ignore

package main

//type CQueue struct {
//	stack1 []int
//	stack2 []int
//}
//
//func Constructor() CQueue {
//	return CQueue{}
//}
//
//func (q *CQueue) AppendTail(value int) {
//	q.stack1 = append(q.stack1, value)
//}
//
//func (q *CQueue) DeleteHead() int {
//	t := -1
//	if len(q.stack2) != 0 {
//		t, q.stack2 = q.stack2[len(q.stack2)-1], q.stack2[:len(q.stack2)-1]
//		return t
//	}
//	if len(q.stack1) != 0 {
//		l := len(q.stack1)
//		q.stack2 = make([]int, l-1)
//		for i := l - 1; i > 0; i-- {
//			q.stack2[l-i-1] = q.stack1[i]
//		}
//		t, q.stack1 = q.stack1[0], q.stack1[:0]
//		return t
//	}
//	return -1
//}

type Stack []int

func (s *Stack) Append(v int) {
	*s = append(*s, v)
}

func (s *Stack) HasTop() bool {
	return s.Size() >= 1
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Pop() (v int) {
	*s, v = (*s)[:len(*s)-1], (*s)[len(*s)-1]
	return
}

type CQueue struct {
	back  Stack
	front Stack
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.back.Append(value)
}

func (this *CQueue) DeleteHead() int {
	if !this.front.HasTop() && !this.back.HasTop() {
		return -1
	}
	if this.front.HasTop() {
		return this.front.Pop()
	}
	for this.back.Size() > 0 {
		this.front.Append(this.back.Pop())
	}
	return this.DeleteHead()
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
