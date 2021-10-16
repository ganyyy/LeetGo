package main

import (
	. "leetgo/data"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历+计数的形式处理吧...

	var tmp []*TreeNode
	var cur = root

	for cur != nil || len(tmp) != 0 {
		if cur != nil {
			tmp = append(tmp, cur)
			cur = cur.Left
		} else {
			cur = tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
			k--
			if k == 0 {
				return cur.Val
			}
			cur = cur.Right
		}
	}
	return math.MinInt32
}
