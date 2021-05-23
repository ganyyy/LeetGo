package main

import (
	"math"
	"sort"
)

// 数字的最大值限定了最大的位数
const L = 30

type nodeNormal struct {
	next [2]*nodeNormal
}

func (n *nodeNormal) insert(v int) {
	for i := L - 1; i >= 0; i-- {
		var idx = (v >> i) & 1
		if n.next[idx] == nil {
			n.next[idx] = newNode()
		}
		n = n.next[idx]
	}
}

func (n *nodeNormal) getMaxXor(val int) (ret int) {
	var t = n
	// 树中的值都是小于限定条件的, 所以这里只需要关注如何取得最大的异或结果即可
	for i := L - 1; i >= 0; i-- {
		var idx = (val >> i) & 1
		if t.next[idx^1] != nil {
			ret |= 1 << i
			idx ^= 1
		}
		t = t.next[idx]
	}
	return
}

func newNode() *nodeNormal {
	return &nodeNormal{}
}

func maximizeXor(nums []int, queries [][]int) []int {
	// 暴力解法...
	// 不可取吗?
	var root = newNode()
	// 前缀树呗...

	// 数字排序
	sort.Ints(nums)
	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	// 查询排序
	sort.Slice(queries, func(i, j int) bool { return queries[i][1] < queries[j][1] })

	var ans = make([]int, len(queries))

	var idx int
	var n = len(nums)
	for _, q := range queries {
		x, m, qid := q[0], q[1], q[2]

		for idx < n && nums[idx] <= m {
			root.insert(nums[idx])
			idx++
		}

		if idx == 0 {
			ans[qid] = -1
		} else {
			ans[qid] = root.getMaxXor(x)
		}
	}

	return ans

}

type trie struct {
	children [2]*trie
	min      int
}

func (t *trie) insert(val int) {
	node := t
	if val < node.min {
		node.min = val
	}
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		if node.children[bit] == nil {
			node.children[bit] = &trie{min: val}
		}
		node = node.children[bit]
		if val < node.min {
			node.min = val
		}
	}
}

func (t *trie) getMaxXorWithLimit(val, limit int) (ans int) {
	node := t
	if node.min > limit {
		return -1
	}
	for i := L - 1; i >= 0; i-- {
		bit := val >> i & 1
		// 每个位置上存放的是最小值, 如果最小值依旧大于限定值, 说明这条分支上的数都大于限定值
		// 这个分支就不能走. 通过这个方法避免了排序
		if node.children[bit^1] != nil && node.children[bit^1].min <= limit {
			ans |= 1 << i
			bit ^= 1
		}
		node = node.children[bit]
	}
	return
}

func maximizeXorWitLimit(nums []int, queries [][]int) []int {
	t := &trie{min: math.MaxInt32}
	for _, val := range nums {
		t.insert(val)
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = t.getMaxXorWithLimit(q[0], q[1])
	}
	return ans
}
