package main

import (
	. "leetgo/data"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	var res = math.MaxInt32
	var pre = math.MaxInt32

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != math.MaxInt32 {
			if t := root.Val - pre; t < res {
				res = t
			}
		}
		pre = root.Val
		dfs(root.Right)
	}
	dfs(root)

	return res
}

// 循环版
func minDiffInBSTLoop(root *TreeNode) int {
	var res = math.MaxInt32
	var pre = math.MaxInt32

	var stack []*TreeNode

	for root != nil || len(stack) != 0 {
		if root == nil {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != math.MaxInt32 {
				res = min(res, root.Val-pre)
			}
			pre = root.Val
			root = root.Right
		} else {
			stack = append(stack, root)
			root = root.Left
		}
	}
	return res
}
