package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

import . "leetgo/data"

func findTilt(root *TreeNode) int {
	var ret int

	var helper func(root *TreeNode) int

	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		var left = helper(root.Left)
		var right = helper(root.Right)
		ret += Abs(left - right)
		return root.Val + left + right
	}

	helper(root)

	return ret
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
