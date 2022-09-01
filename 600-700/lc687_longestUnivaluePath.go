package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(root *TreeNode, p int) int

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var ret int
	dfs = func(root *TreeNode, p int) int {
		if root == nil {
			return 0
		}
		l := dfs(root.Left, root.Val)
		r := dfs(root.Right, root.Val)

		ret = max(ret, l+r)

		if p == root.Val {
			return max(l, r) + 1
		}
		return 0
	}

	dfs(root, root.Val)

	return ret
}
