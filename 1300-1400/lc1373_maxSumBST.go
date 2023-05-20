//go:build ignore

package main

import "math"

func maxSumBST(root *TreeNode) (ans int) {
	// 返回子树的最小值, 子树的最大值, 以及子树的和
	var dfs func(*TreeNode) (int, int, int)
	dfs = func(node *TreeNode) (int, int, int) {
		if node == nil {
			return math.MaxInt, math.MinInt, 0
		}
		// 后序迭代, 由下往上判断
		lMin, lMax, lSum := dfs(node.Left)  // 递归左子树
		rMin, rMax, rSum := dfs(node.Right) // 递归右子树
		x := node.Val
		// 不能像等!
		if x <= lMax || x >= rMin { // 不是二叉搜索树
			return math.MinInt, math.MaxInt, 0
		}
		// 左右都是合法的二叉搜索树, 直接累加和
		s := lSum + rSum + x
		ans = max(ans, s)

		return min(lMin, x), max(rMax, x), s
	}
	dfs(root)
	return
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
