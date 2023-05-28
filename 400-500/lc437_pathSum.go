//go:build ignore

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

	// 用于存储前缀和, 以及前缀和出现的次数
	var count = make(map[int]int)
	// 0代表着这就是一个相加等于 targetSum 的路径
	count[0] = 1

	// 前缀和
	var dfs func(root *TreeNode, pre int) int
	// 避免了回溯过多的问题
	dfs = func(root *TreeNode, pre int) int {
		if root == nil {
			return 0
		}
		var res int
		pre += root.Val
		// 还是区间和. 本质上就是 sum[:i] - sum[:j] == target, 即为 sum[i:j] = target
		// 所以我们只需要找到 sum[:j] = sum[:i] - target 即可
		res += count[pre-targetSum]
		// 更新前缀和
		count[pre]++
		// 递归左右子树
		res += dfs(root.Left, pre) + dfs(root.Right, pre)
		// 回溯
		count[pre]--

		return res
	}

	return dfs(root, 0)
}
