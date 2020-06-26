package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	pre := &ListNode{Next: head}
	var empty = struct{}{}
	m := make(map[int]struct{})
	root := pre
	for head != nil {
		if _, ok := m[head.Val]; !ok {
			m[head.Val] = empty
			pre.Next = head
			pre = pre.Next
		}
		head = head.Next
	}
	pre.Next = nil
	return root.Next
}

func main() {
	l := &ListNode{Val: 1}
	l.Add(2)
	l.Add(3)
	l.Add(3)
	l.Add(2)
	l.Add(1)

	removeDuplicateNodes(l)

	ShowList(l)
}
