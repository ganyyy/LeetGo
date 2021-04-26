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
func rangeSumBST(root *TreeNode, low int, high int) int {
	var res int
	var stack []*TreeNode

	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root.Val > high {
				return res
			}
			if root.Val >= low {
				res += root.Val
			}
			root = root.Right
		}
	}
	return res
}

func rangeSumBST2(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if v := root.Val; v > high {
		return rangeSumBST(root.Left, low, high)
	} else if v < low {
		return rangeSumBST(root.Right, low, high)
	} else {
		return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
	}
}
