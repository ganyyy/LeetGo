//go:build ignore

package main

import (
	. "leetgo/data"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ret int
	var dfs func(root *TreeNode, count int)

	dfs = func(root *TreeNode, count int) {
		if root == nil {
			return
		}
		count ^= (1 << root.Val)
		// 如果是叶子节点?
		if root.Left == nil && root.Right == nil {
			if count == 0 || count&(count-1) == 0 {
				ret++
			}
		} else {
			dfs(root.Left, count)
			dfs(root.Right, count)
		}
	}
	dfs(root, 0)
	return ret
}
