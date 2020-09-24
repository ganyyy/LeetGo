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
func buildTree(inorder []int, postorder []int) *TreeNode {
	// 后序遍历的最后一个是根节点

	// 从中序中分成两部分

	ln := len(inorder)
	if ln == 0 {
		return nil
	}
	last := postorder[ln-1]
	postorder = postorder[:ln-1]
	var pos int
	for pos = range inorder {
		if inorder[pos] == last {
			break
		}
	}
	node := &TreeNode{Val: last}
	node.Left = buildTree(inorder[:pos], postorder[:pos])
	node.Right = buildTree(inorder[pos+1:], postorder[pos:])
	return node
}
