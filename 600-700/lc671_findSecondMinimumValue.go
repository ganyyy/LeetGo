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

	if root == nil {
		return -1
	}
	var val = root.Val
	var dfs func(root *TreeNode) int
	// 根节点一定是最小的!
	dfs = func(root *TreeNode) int {
		if root == nil {
			return -1
		}
		// 那么第一个大于根节点的数一定是第二小的
		if root.Val > val {
			return root.Val
		}

		var left = dfs(root.Left)
		var right = dfs(root.Right)

		// 小于0说明左边/右边到头了
		if left < 0 {
			return right
		}
		if right < 0 {
			return left
		}

		// 否则取两边的最小值
		if right > left {
			return left
		} else {
			return right
		}
	}

	return dfs(root)
}
