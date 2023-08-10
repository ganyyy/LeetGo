package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil && root.Left != nil {
		return 1 + minDepth(root.Left)
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
