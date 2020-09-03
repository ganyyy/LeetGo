package main

import "strconv"
import . "leetgo/data"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	lRes := binaryTreePaths(root.Left)
	rRes := binaryTreePaths(root.Right)
	res := make([]string, 0, len(lRes)+len(rRes))
	cur := strconv.Itoa(root.Val)

	if lRes == nil && rRes == nil {
		// 叶子节点
		res = append(res, cur)
	} else {
		// 普通节点
		cur = cur + "->"
		for _, v := range lRes {
			res = append(res, cur+v)
		}
		for _, v := range rRes {
			res = append(res, cur+v)
		}
	}

	return res
}
