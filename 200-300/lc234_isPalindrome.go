package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	var fast, slow = head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转后半段
	reverseList := reverse(slow.Next)

	// 重头开始对比
	for head != nil && reverseList != nil {
		if head.Val != reverseList.Val {
			return false
		}
		head = head.Next
		reverseList = reverseList.Next
	}
	return true

}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var next, pre *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
