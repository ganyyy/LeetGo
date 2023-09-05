//go:build ignore

package main

import . "leetgo/data"

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var maxDepth int
	var node *TreeNode

	var traverseTree func(root *TreeNode, depth int) int
	traverseTree = func(root *TreeNode, depth int) int {
		if root == nil {
			return depth
		}

		depth++
		l, r := traverseTree(root.Left, depth), traverseTree(root.Right, depth)
		// l和r相等时, 说明当前节点是最近公共祖先, 但不一定是最深的
		if l == r && l >= maxDepth {
			// 为啥要带 = 呢, 假设node是当前节点的左右子节点中的一个, 那么max(l, r)一定是maxDepth
			// 如果恰好当前节点的左右子节点最深的深度相同, 并且 == maxDepth, 那么当前节点就是最近公共祖先
			maxDepth = l
			node = root
		}
		return max(l, r)
	}

	traverseTree(root, 0)

	return node
}
