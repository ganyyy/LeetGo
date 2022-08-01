//go:build ignore

package main

type MyCircularQueue struct {
	buff       []int
	head, tail int
	size       int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		buff: make([]int, k),
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.buff[q.tail] = value
	q.tail = (q.tail + 1) % len(q.buff)
	q.size++
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.head = (q.head + 1) % len(q.buff)
	q.size--
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.buff[q.head]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.buff[(q.tail-1+len(q.buff))%len(q.buff)]
}

func (q *MyCircularQueue) IsEmpty() bool {
	// fmt.Println("IsEmpty", q.buff, q.head, q.tail, q.size)
	return q.size == 0
}

func (q *MyCircularQueue) IsFull() bool {
	// fmt.Println("IsFull", q.buff, q.head, q.tail, q.size)
	return q.size == len(q.buff)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
