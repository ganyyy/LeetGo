package main

type FreqStack struct {
	top   int
	freq  map[int]int   // 每个数字对应的频率
	stack map[int][]int // 相同频率对应的数字队列
}

func Constructor() FreqStack {
	return FreqStack{
		top:   0,
		freq:  make(map[int]int),
		stack: make(map[int][]int),
	}
}

func (s *FreqStack) Push(val int) {
	nf := s.freq[val] + 1
	s.freq[val] = nf
	s.stack[nf] = append(s.stack[nf], val)
	if nf > s.top {
		s.top = nf
	}
	// fmt.Println("Push",s)
}

func (s *FreqStack) Pop() int {
	if s.top == 0 {
		// ? 栈空了
		return 0
	}
	queue := s.stack[s.top]
	if len(queue) == 0 {
		return 0 // 栈空了...
	}
	// 取出队列的末尾元素
	var val = queue[len(queue)-1]
	// 这里, copy更好点?
	queue = queue[:len(queue)-1]
	// 判断队列是否为空
	if len(queue) > 0 {
		// 如果这一层还有元素, 就更新一下
		s.stack[s.top] = queue
	} else {
		// 不删除, 内存无法释放切片内存
		delete(s.stack, s.top)
		// 如果这一层没元素了, 就低一层
		s.top--
	}
	// 更新一下val对应的频率
	old := s.freq[val]
	if old > 1 {
		s.freq[val] = old - 1
	} else {
		delete(s.freq, val)
	}

	// fmt.Println("Pop",s)

	return val
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
