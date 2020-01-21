package main

import (
	"fmt"
	. "leetgo/data"
)

func swapPairs(head *ListNode) *ListNode {
	if nil == head {
		return head
	}

	h := &ListNode{
		Val:  0,
		Next: head,
	}
	i := 0
	ppre := h
	pre := ppre.Next
	head = head.Next
	for nil != head {
		if i & 1 == 0 {
			ppre.Next = head
			pre.Next = head.Next
			head.Next = pre
			head = pre.Next
		} else {
			head = head.Next
		}
		ppre = ppre.Next
		pre = ppre.Next
		i++
	}
	return h.Next
}

func swapPairs2(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	next := head.Next
	head.Next = swapPairs2(next.Next)
	next.Next = head
	return next
}

func main() {
	l := &ListNode{
		Val:  1,
		Next: nil,
	}
	l.Add(2)
	l.Add(3)
	//l.Add(4)

	ShowList(l)
	fmt.Println()
	l = swapPairs2(l)
	ShowList(l)
}


