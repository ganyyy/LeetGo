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
