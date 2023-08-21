package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	// 需要知道中点, 然后对折, 递归处理
	if nil == head {
		return nil
	}
	if nil == head.Next {
		return &TreeNode{Val: head.Val}
	}
	pre := head
	// 一个走两步, 一个走一步, 快指针走到终点慢指针走到中点
	p := pre.Next
	q := p.Next

	for q != nil && q.Next != nil {
		pre = pre.Next
		p = pre.Next
		q = q.Next.Next
	}

	pre.Next = nil
	root := &TreeNode{Val: p.Val}
	// 左子树
	root.Left = sortedListToBST(head)
	// 右子树
	root.Right = sortedListToBST(p.Next)

	return root
}
