package main

import . "leetgo/data"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isEvenOddTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var layer = 1
	var stack []*TreeNode
	stack = append(stack, root)

	for len(stack) != 0 {
		var ln = len(stack)
		var pre = -1
		for i := 0; i < ln; i++ {
			var node = stack[i]
			if node.Val&1 != layer&1 {
				return false
			}
			if pre != -1 {
				if layer&1 == 1 {
					if pre >= node.Val {
						return false
					}
				} else {
					if pre <= node.Val {
						return false
					}
				}
			}

			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			pre = node.Val
		}
		copy(stack, stack[ln:])
		stack = stack[:len(stack)-ln]
		layer++
	}

	return true
}
