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
func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val == t2.Val {
		return checkSubTree(t1.Right, t2.Right) && checkSubTree(t1.Left, t2.Left)
	} else {
		return checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
	}
}

func main() {

}
