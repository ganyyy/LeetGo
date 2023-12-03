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
func bstToGst(root *TreeNode) *TreeNode {
	// 逆中序迭代
	var dfs func(root *TreeNode, pre int) int
	dfs = func(root *TreeNode, pre int) int {
		if root == nil {
			return pre
		}
		root.Val += dfs(root.Right, pre)
		return dfs(root.Left, root.Val)
	}
	dfs(root, 0)
	return root
}
