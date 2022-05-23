package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(*TreeNode, int) bool

	dfs = func(root *TreeNode, v int) bool {
		if root == nil {
			return true
		}
		if root.Val != v {
			return false
		}
		return dfs(root.Left, v) && dfs(root.Right, v)
	}

	return dfs(root, root.Val)
}
