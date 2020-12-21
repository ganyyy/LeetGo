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
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 队列走起
	var queue = make([]*TreeNode, 1, 10)
	queue[0] = root

	var res [][]int
	var t int
	var tmp []int
	var cur *TreeNode
	var layer, begin = 1, 0

	for len(queue) != 0 {
		t = len(queue)
		tmp = make([]int, t)
		if layer > 0 {
			begin = 0
		} else {
			begin = t - 1
		}
		for i := 0; i < t; i++ {
			cur = queue[i]
			tmp[begin] = cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			begin += layer
		}

		queue = queue[t:]
		layer *= -1
		res = append(res, tmp)
	}

	return res
}
