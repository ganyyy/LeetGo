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
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if nil == s {
		return false
	}
	// 当前匹配 或者 左匹配 或者 右匹配
	return isSameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

// 检查当前两个树是否匹配
func isSameTree(s, t *TreeNode) bool {
	if nil == s && nil == t {
		return true
	}
	if nil == s || nil == t {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}

func main() {

}
