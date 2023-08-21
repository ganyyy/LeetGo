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
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// 快慢指针求中点, 然后后半段逆序, 交替插入
	var slow, fast = head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转后半段
	var reverseList = reverse(slow.Next)
	// 清空中点的下一个节点
	slow.Next = nil

	// 从头开始插入
	var cur = head

	var reverseNext, curNext *ListNode
	for cur != nil && reverseList != nil {
		reverseNext = reverseList
		reverseList = reverseList.Next

		curNext = cur.Next
		// 插入
		reverseNext.Next = curNext
		cur.Next = reverseNext

		cur = curNext
	}
}

// 单链表的反转
func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func main() {
	var l = &ListNode{Val: 1}
	l.Add(2)
	l.Add(3)
	l.Add(4)

	reorderList(l)

	ShowList(l)
}
