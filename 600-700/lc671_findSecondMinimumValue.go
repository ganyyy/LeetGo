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
func findSecondMinimumValue(root *TreeNode) int {
	var dfs func(root *TreeNode, val int) int
	// 根节点一定是最小的!
	dfs = func(root *TreeNode, val int) int {
		if root == nil {
			return -1
		}
		// 那么第一个大于根节点的数一定是第二小的
		if root.Val > val {
			return root.Val
		}

		var left = dfs(root.Left, val)
		var right = dfs(root.Right, val)

		if left < 0 {
			return right
		}
		if right < 0 {
			return left
		}
		if right > left {
			return left
		} else {
			return right
		}
	}

	return dfs(root, root.Val)
}
