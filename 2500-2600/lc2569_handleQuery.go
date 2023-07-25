package main

import "math/bits"

type Segments []Segment

type Segment struct {
	left, right int32 // 这个节点所管理的原始数组的子区间
	cnt1        int32 // 区间内 1 的个数
	flip        bool  // 是否反转
}

// 维护区间 1 的个数, 左右子节点的 1 的个数之和
func (s Segments) maintain(pos int32) {
	s[pos].cnt1 = s[pos<<1].cnt1 + s[pos<<1|1].cnt1
}

func (s Segments) build(nums []int, pos, left, right int32) {
	// pos从1开始, 并且越来越大

	s[pos].left, s[pos].right = left, right
	if left == right {
		// 这是一个叶子节点, 1 的个数就是 nums[left-1] 的值
		s[pos].cnt1 = int32(nums[left-1])
		return
	}
	// 递归构造左右子节点,

	mid := (left + right) >> 1
	// 子节点的位置是父节点的两倍([left, m]), 或者两倍加一([m+1, right])
	s.build(nums, pos<<1, left, mid)
	s.build(nums, pos<<1|1, mid+1, right)
	// 维护父节点的 1 的个数
	s.maintain(pos)
}

// 执行区间反转
func (s Segments) reverse(pos int32) {
	seg := &s[pos]
	// 反转的时候, 1 的个数就是区间长度减去原来的 1 的个数
	seg.cnt1 = seg.right - seg.left + 1 - seg.cnt1
	seg.flip = !seg.flip
}

func (s Segments) spread(pos int32) {
	// 这一步是为啥呢..?
	/*
	   比如在 update 这个函数的第一个分支中，
	   如果恰好 left,right 包含了某个完整的子区间，那么此时会针对这个完整的区间设置 flip,
	   此时不会再递归的把这个区间所有的子区间都进行翻转。

	   但是，并不是每次更新都是能恰好涵盖一整个完整的区间的，所以一旦某次更新涉及到了这个区间内的某个子区间，
	   此时就需要先通过 spread 将之前的某次更新应用进去，然后会随着递归逐级的应用，直到到达叶子节点，或者一个更小的子区间才停止。
	*/
	if s[pos].flip {
		s.reverse(pos << 1)
		s.reverse(pos<<1 | 1)
		s[pos].flip = false
	}
}

func (s Segments) update(pos, left, right int32) {
	// 递归更新过程中, left 和 right 值全程不变

	if left <= s[pos].left && s[pos].right <= right {
		// 如果当前区间被包含在更新区间内, 直接反转, 不需要递归更新, 设置 flip 标记
		s.reverse(pos)
		return
	}

	// 此时一定是跨区间更新, pos 一定是非叶子节点

	s.spread(pos) // 这一步是为啥呢..? 因为可能存在之前的更新没有被应用到子区间中

	// m: 中点, 判断更新区间在左子区间还是右子区间, 亦或是跨越两个子区间
	// 如果是偶数个节点的话, m优先取左边的
	// 如果是奇数个节点的话, m就是中间的那个节点
	m := (s[pos].left + s[pos].right) >> 1
	// 递归更新左右子节点
	if left <= m {
		s.update(pos<<1, left, right)
	}
	if m < right {
		s.update(pos<<1|1, left, right)
	}
	s.maintain(pos)
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

func handleQuery(nums1, nums2 []int, queries [][]int) (ans []int64) {
	sum := 0
	for _, x := range nums2 {
		sum += x
	}
	// 为啥要乘 4?
	t := make(Segments, alignUp(len(nums1))*2)
	t.build(nums1, 1, 1, int32(len(nums1)))
	for _, q := range queries {
		if q[0] == 1 {
			t.update(1, int32(q[1]+1), int32(q[2]+1))
		} else if q[0] == 2 {
			sum += q[1] * int(t[1].cnt1)
		} else {
			ans = append(ans, int64(sum))
		}
	}
	return
}

func main() {
	handleQuery([]int{1, 0, 1, 0, 1}, []int{1, 1, 1, 1, 1}, [][]int{{1, 3, 4}, {1, 1, 4}, {2, 1, 0}, {3, 0, 0}})
}
