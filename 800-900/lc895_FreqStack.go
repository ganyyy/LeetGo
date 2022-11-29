package main

type FreqStack struct {
	top   int
	freq  map[int]int
	stack map[int][]int
}

func Constructor() FreqStack {
	return FreqStack{
		top:   0,
		freq:  make(map[int]int),
		stack: make(map[int][]int),
	}
}

func (this *FreqStack) Push(val int) {
	nf := this.freq[val] + 1
	this.freq[val] = nf
	this.stack[nf] = append(this.stack[nf], val)
	if nf > this.top {
		this.top = nf
	}
	// fmt.Println("Push",this)
}

func (this *FreqStack) Pop() int {
	if this.top == 0 {
		// ? 栈空了
		return 0
	}
	queue := this.stack[this.top]
	if len(queue) == 0 {
		return 0 // 栈空了...
	}
	// 取出队列的末尾元素
	var val = queue[len(queue)-1]
	queue = queue[:len(queue)-1]
	// 判断队列是否为空
	if len(queue) > 0 {
		// 如果这一层还有元素, 就更新一下
		this.stack[this.top] = queue
	} else {
		delete(this.stack, this.top)
		// 如果这一层没元素了, 就低一层
		this.top--
	}
	// 更新一下val对应的频率
	old := this.freq[val]
	if old > 1 {
		this.freq[val] = old - 1
	} else {
		delete(this.freq, val)
	}

	// fmt.Println("Pop",this)

	return val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
