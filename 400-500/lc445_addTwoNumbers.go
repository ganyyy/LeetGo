package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var v1, v2 []int
	for l1 != nil {
		v1 = append(v1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		v2 = append(v2, l2.Val)
		l2 = l2.Next
	}

	var head *ListNode
	var t int
	for len(v1) != 0 || len(v2) != 0 || t != 0 {
		if l := len(v1); l != 0 {
			t += v1[l-1]
			v1 = v1[:l-1]
		}
		if l := len(v2); l != 0 {
			t += v2[l-1]
			v2 = v2[:l-1]
		}
		var node = &ListNode{}
		node.Val = t % 10
		node.Next = head
		head = node
		t /= 10
	}
	return head
}

func main() {

}
