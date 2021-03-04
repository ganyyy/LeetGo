package main

type MyQueue struct {
	in, out []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() (val int) {
	q.move()
	var outLen = len(q.out)
	val, q.out = q.out[outLen-1], q.out[:outLen-1]
	return
}

func (q *MyQueue) move() {
	var outLen = len(q.out)
	if outLen == 0 {
		var inLen = len(q.in)
		if inLen == 0 {
			return
		}
		outLen = inLen
		for i := len(q.in) - 1; i >= 0; i-- {
			q.out = append(q.out, q.in[i])
		}
		q.in = q.in[:0]
	}
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	q.move()
	return q.out[len(q.out)-1]
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return len(q.out) == 0 && len(q.in) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
