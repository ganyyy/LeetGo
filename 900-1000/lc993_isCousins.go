package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type myNode struct {
	*TreeNode
	parent *TreeNode
}

func isCousins(r *TreeNode, x int, y int) bool {
	// 层次遍历
	var stack = []*myNode{{TreeNode: r}}
	var fa, fb *myNode
	var tmp []*myNode
	var root *myNode
	for len(stack) != 0 {
		fa, fb = nil, nil
		for i := 0; i < len(stack); i++ {
			root = stack[i]
			if root.Val == x {
				fa = root
			} else if root.Val == y {
				fb = root
			}
			if fa != nil && fb != nil {
				if fa.parent == fb.parent {
					return false
				}
				return true
			}
			if root.Left != nil {
				tmp = append(tmp, &myNode{root.Left, root.TreeNode})
			}
			if root.Right != nil {
				tmp = append(tmp, &myNode{root.Right, root.TreeNode})
			}
		}
		// 找到任意一个均可以提前返回
		if fa != nil || fb != nil {
			return false
		}
		stack = make([]*myNode, len(tmp))
		copy(stack, tmp)
		tmp = tmp[:0]
	}

	return false
}

func isCousinsDFS(root *TreeNode, x, y int) (ans bool) {
	depth := 0
	var father *TreeNode
	var dfs func(*TreeNode, *TreeNode, int) bool
	dfs = func(node, fa *TreeNode, d int) bool {
		if node == nil {
			return false
		}
		if node.Val == x || node.Val == y { // 找到 x 或 y
			if depth > 0 { // 之前已找到 x y 其中一个
				ans = depth == d && father != fa
				return true // 表示 x 和 y 都找到
			}
			depth, father = d, fa // 之前没找到，记录信息
		}
		return dfs(node.Left, node, d+1) || dfs(node.Right, node, d+1)
	}
	dfs(root, nil, 1)
	return
}
