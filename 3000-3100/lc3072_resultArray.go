package main

import (
	"slices"
	"sort"
)

type fenwick []int

// 把下标为 i 的元素增加 1
func (f fenwick) add(i int) {
	// 3 -> 4 -> 8
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}

// 返回下标在 [1,i] 的元素之和
func (f fenwick) pre(i int) (res int) {
	// 7 -> 6 -> 4 -> 0
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func resultArray(nums []int) (ans []int) {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	// 去重后的有序数组的长度
	// 为啥要去重呢? 因为题目要求的是 "绝对大于"
	m := len(sorted)

	a := nums[:1]
	b := []int{nums[1]}
	// 线段树中记录的是位置组成的集合
	// 如果直接二分的话, 需要的内存空间/移动的开销会很大.
	t1 := make(fenwick, m+1)
	t2 := make(fenwick, m+1)
	t1.add(sort.SearchInts(sorted, nums[0]) + 1)
	t2.add(sort.SearchInts(sorted, nums[1]) + 1)
	for _, x := range nums[2:] {
		// 找到x出现的位置, 然后添加到树状数组中
		v := sort.SearchInts(sorted, x) + 1
		// len: 当前的长度.
		// pre: 小于等于x的数字的个数
		gc1 := len(a) - t1.pre(v) // greaterCount(a, v)
		gc2 := len(b) - t2.pre(v) // greaterCount(b, v)
		if gc1 > gc2 || gc1 == gc2 && len(a) <= len(b) {
			a = append(a, x)
			t1.add(v)
		} else {
			b = append(b, x)
			t2.add(v)
		}
	}
	return append(a, b...)
}
