//go:build ignore

package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distributeCoins(root *TreeNode) int {
	// DFS

	var step int
	var dfs func(*TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 不管是 多的还是少的, 本质上都会进行整体的位移,
		// 每一层移动的数量就是到当前节点未知, 所有冗余的节点
		cur := dfs(root.Left) + dfs(root.Right) + root.Val - 1
		step += abs(cur)
		return cur
	}
	dfs(root)

	return step
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
