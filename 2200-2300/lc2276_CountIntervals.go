package main

import "github.com/emirpasic/gods/trees/redblacktree"

type CountIntervals struct {
	*redblacktree.Tree
	cnt int // 所有区间长度和
}

func Constructor() CountIntervals {
	return CountIntervals{redblacktree.NewWithIntComparator(), 0}
}

func (t *CountIntervals) Add(left, right int) {
	// 遍历所有被 [left,right] 覆盖到的区间（部分覆盖也算）
	// [node.val, node.key]
	//              [left, right]
	for node, _ := t.Ceiling(left); node != nil && node.Value.(int) <= right; node, _ = t.Ceiling(left) {
		l, r := node.Value.(int), node.Key.(int)
		left = min(left, l)   // 合并后的新区间，其左端点为所有被覆盖的区间的左端点的最小值
		right = max(right, r) // 合并后的新区间，其右端点为所有被覆盖的区间的右端点的最大值
		t.cnt -= r - l + 1
		t.Remove(r)
	}
	t.cnt += right - left + 1
	// key是右区间, val是左区间
	t.Put(right, left) // 所有被覆盖到的区间与 [left,right] 合并成一个新区间
}

func (t *CountIntervals) Count() int { return t.cnt }
