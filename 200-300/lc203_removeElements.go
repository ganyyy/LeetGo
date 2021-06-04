package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	head = &ListNode{Next: head}
	var pre, cur = head, head.Next
	pre.Next = nil
	for cur != nil {
		if cur.Val != val {
			pre.Next = cur
			pre, cur = cur, cur.Next
			pre.Next = nil
		} else {
			cur = cur.Next
		}
	}
	return head.Next
}
