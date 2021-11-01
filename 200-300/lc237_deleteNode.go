//go:build ignore

package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	// var pre = node
	// for node.Next != nil {
	//     node.Val = node.Next.Val
	//     pre = node
	//     node = node.Next
	// }

	// pre.Next = nil

	// 直接跳过可还行
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
