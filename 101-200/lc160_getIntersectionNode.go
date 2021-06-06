package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 逆序找到第一个不相等的节点, 那他的前置节点就是相交的节点

	// 要么就是搞个map, 找相等的.

	//if headA == nil || headB == nil {
	//	return nil
	//}
	//
	//// 单链表的逆序, 怎么搞不用我说了吧
	//var reverse = func(node *ListNode) *ListNode {
	//	var next = node.Next
	//	node.Next = nil
	//	for next != nil {
	//		var tmp = next
	//		next = next.Next
	//		tmp.Next = node
	//		node = tmp
	//	}
	//	return node
	//}
	//
	//headA = reverse(headA)
	//headB = reverse(headB)
	//
	//if headA.Val != headB.Val {
	//	reverse(headA)
	//	reverse(headB)
	//	return nil
	//}
	//
	//var pre *ListNode
	//
	//for curA, curB := headA, headB; curA != nil && curB != nil && curA.Val == curB.Val; curA, curB = curA.Next, curB.Next {
	//	pre = curA
	//}
	//
	//reverse(headA)
	//reverse(headB)
	//
	//return pre

	// ???, 用map搞一下呗...

	// Map 可以直接求解

	// 成吧

	if headA == nil || headB == nil {
		return nil
	}

	// 最多执行 len(head1)+len(head2)+1
	// 学会从条件中找到恒定的数据, 往往这就是答案的解

	var ca, cb = headA, headB
	for ca != cb {
		if ca == nil {
			ca = headB
		} else {
			ca = ca.Next
		}
		if cb == nil {
			cb = headA
		} else {
			cb = cb.Next
		}
	}
	return ca
}

func main() {
	var l1 = &ListNode{Val: 4}
	var l2 = &ListNode{Val: 5}
	var l3 = &ListNode{Val: 8}
	l3.Add(4).Add(5)
	l1.Add(1).Next = l3
	l2.Add(6).Add(1).Next = l3

	ShowList(l1)
	ShowList(l2)
	ShowList(getIntersectionNode(l1, l2))
}
