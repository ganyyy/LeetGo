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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}

	var res [][]int
	for len(queue) != 0 {
		l := len(queue)
		var tmp = make([]int, l)
		for i := 0; i < l; i++ {
			top := queue[i]
			tmp[i] = top.Val
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		res = append(res, tmp)
		queue = queue[l:]
	}
	// 逆序
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return res
}
