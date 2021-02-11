package main

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	k    int
	nums []int

	ia IntArray
}

func Constructor2(k int, nums []int) KthLargest {
	var tmp = nums
	sort.Ints(nums)
	if len(nums) < k {
		tmp = make([]int, len(nums), k)
		copy(tmp, nums)
	} else {
		tmp = tmp[len(nums)-k:]
	}
	return KthLargest{
		k:    k,
		nums: tmp,
	}
}

func (k *KthLargest) Add2(val int) int {
	if len(k.nums) < k.k {
		k.insert(val)
		// 即使插入前不够, 插入后也一定够了
		// 不然不符合题意
		return k.nums[0]
	}

	if val <= k.nums[0] {
		return k.nums[0]
	}
	// 前移, 去尾
	copy(k.nums, k.nums[1:])
	k.nums = k.nums[:len(k.nums)-1]
	// 插入
	k.insert(val)
	// 返回结果
	return k.nums[0]
}

func (k *KthLargest) insert(val int) {
	var l, r = 0, len(k.nums)
	k.nums = append(k.nums, val)
	for l < r {
		var mid = l + (r-l)>>1
		if k.nums[mid] == val {
			l = mid
			break
		} else if k.nums[mid] > val {
			r = mid
		} else {
			l = mid + 1
		}
	}
	// 插进去
	copy(k.nums[l+1:], k.nums[l:])
	k.nums[l] = val
}

// 这个方法不太好, 我们可以用小顶堆进行处理

type IntArray []int

func (ia IntArray) Len() int {
	return len(ia)
}

func (ia IntArray) Less(i, j int) bool {
	return ia[i] < ia[j]
}

func (ia *IntArray) Swap(i, j int) {
	var t = *ia
	t[i], t[j] = t[j], t[i]
	*ia = t
}

func (ia *IntArray) Push(x interface{}) {
	*ia = append(*ia, x.(int))
}

func (ia *IntArray) Pop() (x interface{}) {
	x, *ia = (*ia)[ia.Len()-1], (*ia)[:ia.Len()-1]
	return
}

func Constructor(k int, nums []int) KthLargest {
	var kl = KthLargest{
		k: k,
	}
	if len(nums) < k {
		kl.ia = make([]int, len(nums), k)
		copy(kl.ia, nums)
		heap.Init(&kl.ia)
	} else {
		sort.Ints(nums)
		kl.ia = nums[len(nums)-k:]
	}

	return kl
}

func (k *KthLargest) Add(val int) int {
	if k.ia.Len() < k.k {
		heap.Push(&k.ia, val)
		return k.ia[0]
	}
	if val <= k.ia[0] {
		return k.ia[0]
	}
	heap.Pop(&k.ia)
	heap.Push(&k.ia, val)
	return k.ia[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor2(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	var k = Constructor(3, []int{5, -1})
	println(k.Add(2))
	println(k.Add(1))
	println(k.Add(-1))
	println(k.Add(3))
	println(k.Add(4))
}
