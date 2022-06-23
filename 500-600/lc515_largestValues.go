package main

import (
	"math"

	. "leetgo/data"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var pre, cur []*TreeNode
	var ret []int

	pre = append(pre, root)

	for len(pre) != 0 {
		var mm = math.MinInt32
		for _, node := range pre {
			if node.Val > mm {
				mm = node.Val
			}
			if node.Left != nil {
				cur = append(cur, node.Left)
			}
			if node.Right != nil {
				cur = append(cur, node.Right)
			}
		}
		ret = append(ret, mm)
		pre, cur = cur, pre[:0]
	}

	return ret
}
