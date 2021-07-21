package main

import . "leetgo/data"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var hA, hB = headA, headB
	for headA != nil || headB != nil {
		if headA == headB {
			return headA
		}
		if headA == nil {
			headA = hB
		} else {
			headA = headA.Next
		}
		if headB == nil {
			headB = hA
		} else {
			headB = headB.Next
		}
	}
	return nil
}
