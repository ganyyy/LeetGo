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
func minCameraCover(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var cnt int
	var helper func(root *TreeNode) int
	// 每一个节点有三个状态:
	// 0. 未被照射到
	// 1. 被照射到
	// 2. 自己安装了监视器
	helper = func(root *TreeNode) int {
		// 为空的节点可以理解为被照射到
		if root == nil {
			return 1
		}
		// 看左右两个子树的状态
		left, right := helper(root.Left), helper(root.Right)
		// 如果有一个未被照射到, 那么就在自己身上安装一个监视器
		if left == 0 || right == 0 {
			cnt++
			return 2
		} else if left == 2 || right == 2 {
			// 左右子树有一个安装了监视器, 那么自己就属于被照射到的那个
			return 1
		} else {
			// 自己未安装, 通知父节点处理
			return 0
		}
	}

	// 如果根节点需要进行处理, 只能在自己身上加一个监视器
	if helper(root) == 0 {
		cnt++
	}

	return cnt
}
