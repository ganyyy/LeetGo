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
func hasPathSum(root *TreeNode, sum int) bool {

	if nil == root {
		return false
	}

	remain := sum - root.Val
	if nil == root.Left && nil == root.Right {
		// 叶子节点并且 看看此时值和目标值是否相等
		return remain == 0
	} else {
		// 返回左子树或者右子树的判断
		return nil != root.Left && hasPathSum(root.Left, remain) || nil != root.Right && hasPathSum(root.Right, remain)
	}
}
