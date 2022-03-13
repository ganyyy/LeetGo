package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(A *TreeNode, B *TreeNode) bool {
		if A == nil && B == nil {
			return true
		}
		if A == nil || B == nil {
			return false
		}
		if A.Val != B.Val {
			return false
		}
		return dfs(A.Left, B.Right) && dfs(A.Right, B.Left)
	}
	return dfs(root.Left, root.Right)
}
