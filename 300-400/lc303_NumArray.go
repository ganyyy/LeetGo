package main

type NumArray struct {
	sum []int
}

func Constructor303(nums []int) NumArray {
	var tmp = make([]int, len(nums))
	var cur int
	for i, v := range nums {
		cur += v
		tmp[i] = cur
	}
	return NumArray{sum: tmp}
}

func (n *NumArray) SumRange(i int, j int) int {
	var sub int
	if i-1 >= 0 {
		sub = n.sum[i-1]
	}
	return n.sum[j] - sub
}

/**
 * Your NumArray307 object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
