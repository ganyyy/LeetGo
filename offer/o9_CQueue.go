package main

type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (q *CQueue) AppendTail(value int) {
	q.stack1 = append(q.stack1, value)
}

func (q *CQueue) DeleteHead() int {
	t := -1
	if len(q.stack2) != 0 {
		t, q.stack2 = q.stack2[len(q.stack2)-1], q.stack2[:len(q.stack2)-1]
		return t
	}
	if len(q.stack1) != 0 {
		l := len(q.stack1)
		q.stack2 = make([]int, l-1)
		for i := l - 1; i > 0; i-- {
			q.stack2[l-i-1] = q.stack1[i]
		}
		t, q.stack1 = q.stack1[0], q.stack1[:0]
		return t
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
