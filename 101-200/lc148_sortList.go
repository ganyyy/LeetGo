package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	// merge sort
	return mergeSort(head)
}

// merge sort
func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 双指针找到中点, slow的下一个就是分界点
	var slow, fast = head, head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var r = mergeSort(slow.Next)
	// 这里需要断开, 不然会内存溢出
	slow.Next = nil
	var l = mergeSort(head)

	return mergeList(l, r)
}

func mergeList(l, r *ListNode) *ListNode {
	var tmp = &ListNode{}
	var p = tmp
	for l != nil && r != nil {
		if l.Val < r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}
	if l == nil {
		p.Next = r
	} else {
		p.Next = l
	}
	return tmp.Next
}
