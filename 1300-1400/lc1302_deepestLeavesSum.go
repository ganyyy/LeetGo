package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	var dfs func(node *TreeNode, depth int)
	ans, mx := 0, 0
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth > mx {
			mx, ans = depth, node.Val
		} else {
			if depth == mx {
				ans += node.Val
			}
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return ans
}
