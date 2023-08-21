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
func isBalanced(root *TreeNode) bool {

	var height func(*TreeNode) int

	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := height(root.Left)
		r := height(root.Right)
		if l == -1 || r == -1 {
			return -1
		}
		sub, max := abs(l, r)
		if sub > 1 {
			return -1
		} else {
			return max + 1
		}
	}

	return height(root) != -1
}

func abs(a, b int) (int, int) {
	if a < b {
		return b - a, b
	} else {
		return a - b, a
	}
}
