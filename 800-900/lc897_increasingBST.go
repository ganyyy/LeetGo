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
func increasingBST(root *TreeNode) *TreeNode {
	var cur, head *TreeNode
	var tmp []*TreeNode
	for root != nil || len(tmp) != 0 {
		if root != nil {
			tmp = append(tmp, root)
			root = root.Left
		} else {
			root = tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
			if head == nil {
				head = root
				cur = head
			} else {
				cur.Right = root
				cur = cur.Right
			}
			root.Left = nil
			root = root.Right
		}
	}

	return head
}

func increasingBST2(root *TreeNode, tail ...*TreeNode) *TreeNode {
	if root == nil {
		if len(tail) == 0 {
			return nil
		} else {
			return tail[0]
		}
	}
	// tail 表示的是当前节点的后继节点

	// 对于左孩子而言, 它的后继节点就是 root
	var left = increasingBST2(root.Left, root)
	root.Left = nil

	// 对于右孩子而言, 它的后继节点是 root 的后继节点
	// root, root.L, root.R, tail之间的关系为
	// root.L < root < root.R < tail
	root.Right = increasingBST2(root.Right, tail...)

	return left
}
