package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var smallArray = &ListNode{}
	var bigArray = &ListNode{}
	smallHead := smallArray
	bigHead := bigArray
	for head != nil {
		if head.Val >= x {
			bigArray.Next = head
			bigArray = head
		} else {
			smallArray.Next = head
			smallArray = head
		}
		head = head.Next
	}
	bigArray.Next = nil
	smallArray.Next = bigHead.Next
	return smallHead.Next
}
