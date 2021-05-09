package main

import . "leetgo/data"

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	// 中序遍历树, 获取其叶子节点

	var stack []*TreeNode
	var val []int

	for len(stack) != 0 || root1 != nil {
		if root1 != nil {
			stack = append(stack, root1)
			root1 = root1.Left
		} else {
			root1 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root1.Left == nil && root1.Right == nil {
				val = append(val, root1.Val)
			}
			root1 = root1.Right
		}
	}

	stack = stack[:0]
	var idx = 0

	root1 = root2
	for len(stack) != 0 || root1 != nil {
		if root1 != nil {
			stack = append(stack, root1)
			root1 = root1.Left
		} else {
			root1 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root1.Left == nil && root1.Right == nil {
				if idx >= len(val) {
					return false
				}
				if root1.Val != val[idx] {
					return false
				}
				idx++
			}
			root1 = root1.Right
		}
	}

	return idx == len(val)
}
