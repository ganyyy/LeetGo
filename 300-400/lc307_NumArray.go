package main

type NumArray struct {
	nums, tree []int
}

// 这是树状数组的解法
func lowBit(x int) int {
	return x & (-x)
}

func (na *NumArray) add(index, val int) {
	// 从低到高更新
	// 3 -> 4 -> 8
	// 1 -> 2 -> 4 -> 8
	for ; index < len(na.tree); index += lowBit(index) {
		na.tree[index] += val
	}
}

func (na *NumArray) prefixSum(index int) int {
	var sum int
	// 从高到低累加
	// 8 -> 0
	// 7 -> 6 -> 4 -> 0
	for ; index > 0; index &= index - 1 {
		sum += na.tree[index]
	}
	return sum
}

func Constructor(nums []int) NumArray {
	var tree = make([]int, len(nums)+1)
	var arr = NumArray{
		nums: nums,
		tree: tree,
	}
	for i := 1; i < len(nums)+1; i++ {
		arr.add(i, nums[i-1])
	}
	return arr
}

func (na *NumArray) Update(index int, val int) {
	na.add(index+1, val-na.nums[index])
	na.nums[index] = val
}

func (na *NumArray) SumRange(left int, right int) int {
	return na.prefixSum(right+1) - na.prefixSum(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
