package main

import (
	"math/bits"
	"sort"
)

type segmentItem struct {
	candidate int
	count     int
}

// MajorityChecker define
type MajorityChecker struct {
	segmentTree []segmentItem
	data        []int
	count       map[int][]int
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

// Constructor define
func Constructor(arr []int) MajorityChecker {
	// data: 原始的数据
	var data = make([]int, len(arr))
	// count: 对应数字的所有下标
	var count = make(map[int][]int)
	// tree: 后半部分是根节点, 前半部分是父节点. 因为数组长度不总是 2的整数次幂, 所以这里至少要 x 4(?)
	//       这里理论上 会趋向于2的整数次幂进行逼近, 然后再 *2
	// 		 严格意义上来讲, 会出现空位
	var tree = make([]segmentItem, alignUp(len(arr))*2)

	var mc MajorityChecker
	for i := 0; i < len(arr); i++ {
		data[i] = arr[i]
	}
	for i := 0; i < len(arr); i++ {
		if _, ok := count[arr[i]]; !ok {
			count[arr[i]] = []int{}
		}
		count[arr[i]] = append(count[arr[i]], i)
	}
	mc.data, mc.segmentTree, mc.count = data, tree, count
	if len(arr) > 0 {
		mc.buildSegmentTree(0, 0, len(arr)-1)
	}
	return mc
}

// 这个 merge 函数就是摩尔投票算法
func (mc *MajorityChecker) merge(i, j segmentItem) segmentItem {
	if i.candidate == j.candidate {
		return segmentItem{candidate: i.candidate, count: i.count + j.count}
	}
	if i.count > j.count {
		return segmentItem{candidate: i.candidate, count: i.count - j.count}
	}
	return segmentItem{candidate: j.candidate, count: j.count - i.count}
}

func (mc *MajorityChecker) buildSegmentTree(treeIndex, left, right int) {
	if left == right {
		mc.segmentTree[treeIndex] = segmentItem{candidate: mc.data[left], count: 1}
		return
	}
	// 获取左右孩子节点
	leftTreeIndex := mc.leftChild(treeIndex)
	rightTreeIndex := mc.rightChild(treeIndex)
	// 获取分界点
	midTreeIndex := left + (right-left)>>1
	mc.buildSegmentTree(leftTreeIndex, left, midTreeIndex)
	mc.buildSegmentTree(rightTreeIndex, midTreeIndex+1, right)
	// 摩尔投票具有累加性
	mc.segmentTree[treeIndex] = mc.merge(mc.segmentTree[leftTreeIndex], mc.segmentTree[rightTreeIndex])
}

func (mc *MajorityChecker) leftChild(index int) int  { return 2*index + 1 }
func (mc *MajorityChecker) rightChild(index int) int { return 2*index + 2 }

// Query define
func (mc *MajorityChecker) query(left, right int) segmentItem {
	if len(mc.data) > 0 {
		return mc.queryInTree(0, 0, len(mc.data)-1, left, right)
	}
	return segmentItem{candidate: -1, count: -1}
}

func (mc *MajorityChecker) queryInTree(treeIndex, left, right, queryLeft, queryRight int) segmentItem {
	// treeIndex 会一直增加
	// [left, right] 会逐渐缩小
	// [queryLeft, queryRight] 会在跨区间时逐渐时逐渐缩小

	// 查询的中点
	mid := left + (right-left)>>1
	// 数据区间的 左右子节点
	leftTreeIndex := mc.leftChild(treeIndex)
	rightTreeIndex := mc.rightChild(treeIndex)

	if queryLeft <= left && queryRight >= right { // segment completely inside range
		// 这个可以加速查询, 底层是两两合并的, 如果某个查询区间涵盖了一个较大的 区间, 可以直接返回
		return mc.segmentTree[treeIndex]
	}
	if queryLeft > mid {
		// 查询区间在右子节点
		return mc.queryInTree(rightTreeIndex, mid+1, right, queryLeft, queryRight)
	} else if queryRight <= mid {
		// 查询区间在左字节带你
		return mc.queryInTree(leftTreeIndex, left, mid, queryLeft, queryRight)
	}
	// 查询区间在左子节点和右子节点的中间
	return mc.merge(mc.queryInTree(leftTreeIndex, left, mid, queryLeft, mid),
		mc.queryInTree(rightTreeIndex, mid+1, right, mid+1, queryRight))
}

// Query define
func (mc *MajorityChecker) Query(left int, right int, threshold int) int {
	res := mc.query(left, right)
	// res可能是-1
	candidateIdx, ok := mc.count[res.candidate]
	if !ok {
		return -1
	}
	// 本身筛选出来的就是众数, 只需要判定一下是否满足threshold即可
	start := sort.Search(len(candidateIdx), func(i int) bool { return left <= candidateIdx[i] })
	end := sort.Search(len(candidateIdx), func(i int) bool { return right < candidateIdx[i] }) - 1
	if (end - start + 1) >= threshold {
		return res.candidate
	}
	return -1
}

func main() {
	// Constructor([]int{2, 2, 1, 2, 1, 2, 2, 1, 1, 2})

	println(alignUp(10))
	println(alignUp(20))
	println(alignUp(30))
	println(alignUp(40))
}
