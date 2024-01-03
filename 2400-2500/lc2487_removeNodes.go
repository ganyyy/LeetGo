package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head = reverse(head)

	pre := head
	for node := head.Next; node != nil; node = node.Next {
		if node.Val >= pre.Val {
			pre.Next = node
			pre = node
		}
	}
	pre.Next = nil
	return reverse(head)
}

func removeNodesStack(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var stack []*ListNode
	for head != nil {
		for len(stack) != 0 && stack[len(stack)-1].Val < head.Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, head)
		head = head.Next
	}
	if len(stack) == 0 {
		return nil
	}
	head = stack[0]
	pre := head
	for _, next := range stack[1:] {
		pre.Next = next
		pre = next
	}
	pre.Next = nil
	return head
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var root ListNode
	for head != nil {
		next := head.Next
		head.Next = root.Next
		root.Next = head
		head = next
	}
	return root.Next

}
