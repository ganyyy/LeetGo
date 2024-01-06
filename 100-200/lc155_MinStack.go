package main

type MinStack struct {
	val []int // 数值的入栈顺序
	min []int // 最小值的入栈顺序
}

func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(val int) {
	ln := len(ms.val) - 1
	ms.val = append(ms.val, val)
	if ln >= 0 && val > ms.min[ln] {
		val = ms.min[ln]
	}
	ms.min = append(ms.min, val)
}

func (ms *MinStack) Pop() {
	ln := len(ms.val)
	if ln == 0 {
		return
	}
	ms.val = ms.val[:ln-1]
	ms.min = ms.min[:ln-1]
}

func (ms *MinStack) Top() int {
	ln := len(ms.val)
	if ln == 0 {
		return 0
	}
	return ms.val[ln-1]
}

func (ms *MinStack) GetMin() int {
	ln := len(ms.val)
	if ln == 0 {
		return 0
	}
	return ms.min[ln-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor146_2();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

}
