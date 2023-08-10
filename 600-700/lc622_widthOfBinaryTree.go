//go:build ignore

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {

	depth := make(map[int]int)

	var dfs func(root *TreeNode, idx int, dep int) int

	dfs = func(root *TreeNode, idx int, dep int) int {
		if root == nil {
			return 0
		}
		if _, ok := depth[dep]; !ok {
			depth[dep] = idx
		}
		return max(idx-depth[dep]+1, max(dfs(root.Left, idx*2, dep+1), dfs(root.Right, idx*2+1, dep+1)))
	}

	return dfs(root, 0, 1)
}
