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
func pathSum(root *TreeNode, targetSum int) int {

	var count = make(map[int]int)
	count[0] = 1

	var helper func(root *TreeNode, pre int) int
	// 避免了回溯过多的问题
	helper = func(root *TreeNode, pre int) int {
		if root == nil {
			return 0
		}
		var res int
		pre += root.Val
		// 还是区间和. 本质上就是 sum[:i] - sum[:j] == target, 即为 sum[i:j] = target
		res += count[pre-targetSum]
		count[pre]++
		res += helper(root.Left, pre) + helper(root.Right, pre)
		count[pre]--

		return res
	}

	return helper(root, 0)
}
