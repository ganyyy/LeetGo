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
func sumOfLeftLeaves(root *TreeNode) int {
	var helper func(root *TreeNode) int

	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		var res int
		if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
			res = root.Left.Val
		}
		return res + helper(root.Left) + helper(root.Right)
	}

	return helper(root)
}
