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
func recoverTree(root *TreeNode) {
	if nil == root {
		return
	}
	// 二叉搜索树的中序遍历是有序的

	// 记录两个有问题的点
	var t1, t2, pre *TreeNode

	var inOrder func(*TreeNode)

	inOrder = func(root *TreeNode) {
		if nil == root {
			return
		}
		inOrder(root.Left)
		// 左边的点应该比右边的点小
		if nil != pre && pre.Val > root.Val {
			if nil == t1 {
				t1 = pre
			}
			// 这里的t2 指的是一种[3,2,1]的情况
			t2 = root
		}
		pre = root
		inOrder(root.Right)
	}

	// 非递归版, 需要一个栈存储临时的变量
	stack := []*TreeNode{}

	for nil != root || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if nil != pre && pre.Val > root.Val {
				if t1 == nil {
					t1 = pre
				}
				t2 = root
			}
			pre = root
			root = root.Right
		}
	}

	if nil == t1 || nil == t2 {
		return
	}
	t1.Val, t2.Val = t2.Val, t1.Val
}
