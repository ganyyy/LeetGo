//go:build ignore

package main

type MyCircularDeque struct {
	queue []int
	head  int
	tail  int
	count int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue: make([]int, k),
		head:  1,
	}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}
	d.head = (d.head - 1 + len(d.queue)) % len(d.queue)
	d.queue[d.head] = value
	d.count++
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}
	d.tail = (d.tail + 1) % len(d.queue)
	d.queue[d.tail] = value
	d.count++
	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	d.head = (d.head + 1) % len(d.queue)
	d.count--
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	d.tail = (d.tail - 1 + len(d.queue)) % len(d.queue)
	d.count--
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.queue[d.head]
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.queue[d.tail]
}

func (d *MyCircularDeque) IsEmpty() bool {
	// fmt.Println("Empty", d.queue)
	return d.count == 0
}

func (d *MyCircularDeque) IsFull() bool {
	// fmt.Println("Full", d.queue)
	return d.count == len(d.queue)
}
