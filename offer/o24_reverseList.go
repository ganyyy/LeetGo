//go:build ignore
// +build ignore

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var root = &ListNode{}
	var cur = root
	for head != nil {
		var next = head.Next
		head.Next = cur.Next
		cur.Next = head
		head = next
	}

	return root.Next
}
