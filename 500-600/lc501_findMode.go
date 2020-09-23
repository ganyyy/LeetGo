package main

import . "leetgo/data"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	var res []int
	// 当前计数以及最大计数
	var cur = 1
	var max int
	var pre *TreeNode
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}
		helper(root.Left)
		if pre != nil {
			if root.Val == pre.Val {
				cur++
			} else {
				cur = 1
			}
		}
		if cur == max {
			res = append(res, root.Val)
		} else if cur > max {
			res = res[:0]
			res = append(res, root.Val)
			max = cur
		}
		pre = root
		helper(root.Right)
	}

	helper(root)
	return res
}
