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
func averageOfLevels(root *TreeNode) []float64 {
	if nil == root {
		return nil
	}
	// 层次遍历
	queue := []*TreeNode{root}
	var res []float64
	for t := len(queue); t != 0; t = len(queue) {
		var sum int
		for i := 0; i < t; i++ {
			c := queue[i]
			sum += c.Val
			if c.Left != nil {
				queue = append(queue, c.Left)
			}
			if c.Right != nil {
				queue = append(queue, c.Right)
			}
		}
		res = append(res, float64(sum)/float64(t))
		queue = queue[t:]
	}

	return res
}
