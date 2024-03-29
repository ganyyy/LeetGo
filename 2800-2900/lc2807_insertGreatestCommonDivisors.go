package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var next *ListNode
	for pre := head; pre != nil && pre.Next != nil; pre = next {
		next = pre.Next
		pre.Next = &ListNode{
			Next: next,
			Val:  gcd(pre.Val, next.Val),
		}
	}
	return head
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
