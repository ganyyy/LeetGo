//go:build ignore

package main

type FrontMiddleBackQueue struct {
	// fl >= bl, 但是相差不能超过1
	// 用链表?
	front, back []int
}

var buffer1 [1000]int
var buffer2 [1000]int

func Constructor1670() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{
		front: buffer1[:0],
		back:  buffer2[:0],
	}
}

func (this *FrontMiddleBackQueue) balance() {
	var balance = func(a, b []int) ([]int, []int) {
		// 将a的第一个元素添加到b的头部
		b = append(b, 0)
		copy(b[1:], b)
		b[0] = a[0]
		a = a[1:]
		return a, b
	}

	fl, bl := len(this.front), len(this.back)
	if fl-bl > 1 {
		// 将front最老的元素移动到back中
		this.front, this.back = balance(this.front, this.back)
	} else if bl-fl > 0 {
		// 将front中最老的元素移动到back中
		this.back, this.front = balance(this.back, this.front)
	}
	// fmt.Println("Front", this.front, "Back", this.back)
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.front = append(this.front, val)
	this.balance()
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	this.front = append(this.front, 0)
	copy(this.front[1:], this.front)
	this.balance()
	this.front[0] = val
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.back = append(this.back, val)
	this.balance()
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.front) < 1 {
		return -1
	}
	fl := len(this.front)
	val := this.front[fl-1]
	this.front = this.front[:fl-1]
	this.balance()
	return val
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.front) < 1 {
		return -1
	}
	val := this.front[0]
	this.front = this.front[1:]
	this.balance()
	return val
}

func (this *FrontMiddleBackQueue) PopBack() int {
	fl, bl := len(this.front), len(this.back)
	if fl == 0 && bl == 0 {
		return -1
	}
	if bl == 0 {
		val := this.front[0]
		this.front = this.front[:0]
		return val
	}
	val := this.back[bl-1]
	this.back = this.back[:bl-1]
	this.balance()
	return val
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
