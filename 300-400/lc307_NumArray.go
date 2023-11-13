package main

import "math/bits"

type NumArray307 struct {
	nums, tree []int
}

// 这是树状数组的解法
func lowBit(x int) int {
	return x & (-x)
}

func (na *NumArray307) add(index, val int) {
	// 从低到高更新
	// 3 -> 4 -> 8
	// 1 -> 2 -> 4 -> 8
	for ; index < len(na.tree); index += lowBit(index) {
		na.tree[index] += val
	}
}

func (na *NumArray307) prefixSum(index int) int {
	var sum int
	// 从高到低累加
	// 8 -> 0
	// 7 -> 6 -> 4 -> 0
	for ; index > 0; index &= index - 1 {
		sum += na.tree[index]
	}
	return sum
}

func Constructor307(nums []int) NumArray307 {
	var tree = make([]int, len(nums)+1)
	var arr = NumArray307{
		nums: nums,
		tree: tree,
	}
	for i := 1; i < len(nums)+1; i++ {
		arr.add(i, nums[i-1])
	}
	return arr
}

func (na *NumArray307) Update(index int, val int) {
	na.add(index+1, val-na.nums[index])
	na.nums[index] = val
}

func (na *NumArray307) SumRange(left int, right int) int {
	return na.prefixSum(right+1) - na.prefixSum(left)
}

/**
 * Your NumArray307 object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */

type NumArray3072 struct {
	// 线段树太操蛋了, 日
	numbers []int
}

func alignUp(n int) int {
	if n <= 0 {
		return 1
	}
	if n&(n-1) == 0 {
		return n
	}
	return 1 << (bits.Len(uint(n)))
}

func Constructor3072(nums []int) NumArray3072 {
	arr := NumArray3072{
		numbers: make([]int, alignUp(len(nums))*2),
	}
	offset := len(arr.numbers) / 2
	copy(arr.numbers[offset:], nums)
	for i := offset - 1; i >= 1; i-- {
		arr.numbers[i] = arr.numbers[i<<1] + arr.numbers[(i<<1)+1]
	}
	return arr
}

func (arr *NumArray3072) Update(index int, val int) {
	offset := len(arr.numbers) / 2
	if index >= offset {
		return
	}
	next := index + offset
	diff := val - arr.numbers[next]
	for ; next >= 1; next >>= 1 {
		arr.numbers[next] += diff
	}
}

func (arr *NumArray3072) preSum(length int) int {
	offset := len(arr.numbers) / 2
	if length > offset {
		length = offset
	}
	var ret int
	for length != 0 {
		count := bits.TrailingZeros64(uint64(length))
		if count == 0 {
			ret += arr.numbers[offset+length-1]
			count = 1
		}
		length >>= count
		offset >>= count
	}
	return ret
}

func (arr *NumArray3072) SumRange(left int, right int) int {
	return arr.preSum(right+1) - arr.preSum(left)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
