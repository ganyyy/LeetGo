package main

type MinStack struct {
	Vals []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	// 插入两个值, 一个是当前元素本身，一个是当前的最小值
	// 空间换时间
	if ln := len(this.Vals); 0 == ln {
		this.Vals = append(this.Vals, x, x)
	} else {
		min := this.Vals[ln-1]
		if min > x {
			min = x
		}
		this.Vals = append(this.Vals, x, min)
	}
}

func (this *MinStack) Pop() {
	if ln := len(this.Vals); 0 != ln {
		this.Vals = this.Vals[:ln-2]
	}
}

func (this *MinStack) Top() int {
	if ln := len(this.Vals); 0 != ln {
		return this.Vals[ln-2]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if ln := len(this.Vals); 0 != ln {
		return this.Vals[ln-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

}
