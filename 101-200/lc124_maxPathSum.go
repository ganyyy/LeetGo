//go:build ignore

package main

import (
	"fmt"
	"math"
)

/*
给定一个非空二叉树，返回其最大路径和。
本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var m = math.MinInt32
	r(root, &m)
	return m
}

func r(root *TreeNode, m *int) int {
	if nil == root {
		return math.MinInt32
	}
	l, r := r(root.Left, m), r(root.Right, m)
	b := root.Val
	// 如果是连续的
	t := max(b, max(b+r, b+l))
	// 如果发生了断层
	*m = max(*m, max(t, max(max(l, r), b+l+r)))
	// fmt.Printf(" base:%d, left:%d, right:%d, *m:%d, 连续:%d, 断层:%d\n", b, l, r, *m, t, max(max(l,r), b+l+r))
	return t
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSumNew(root *TreeNode) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var m = math.MinInt32

	var dfs func(root *TreeNode) int // 以root为根节点, 最大的路径和

	dfs = func(root *TreeNode) int {
		if root == nil {
			return math.MinInt32 // 不能返回0, 因为可能会出现比0小的值
		}
		var left, right = dfs(root.Left), dfs(root.Right)
		var val = root.Val
		var maxSub = max(left, right)
		var cur = max(val, val+maxSub) // 当前根节点的最大路径和

		// 可以是之前的最大值, 也可以是当前值, 也可以是左右子树的最大值, 也可以是根+左+右的值
		m = max(m, max(cur, max(maxSub, val+left+right))) // 累积的最大路径和

		return cur
	}

	dfs(root)

	return m
}

func main() {
	t := &TreeNode{Val: -3}
	fmt.Println(maxPathSum(t))
}
