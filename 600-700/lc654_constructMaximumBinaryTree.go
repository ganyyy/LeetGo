package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mx, mi := math.MinInt32, -1
	for i, v := range nums {
		if v > mx {
			mi = i
			mx = v
		}
	}

	return &TreeNode{
		Val:   mx,
		Left:  constructMaximumBinaryTree(nums[:mi]),
		Right: constructMaximumBinaryTree(nums[mi+1:]),
	}
}
