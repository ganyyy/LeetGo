package main

import . "leetgo/data"

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var h1, h2 = head, head.Next
	var n1, n2 = h1, h2

	for n1.Next != nil && n2.Next != nil {
		n1.Next = n2.Next
		n1 = n1.Next
		n2.Next = n1.Next
		n2 = n2.Next
	}

	n1.Next = h2
	return h1

}

func main() {
	var h = &ListNode{Val: 1}
	//h.Add(2)
	//h.Add(3)
	//h.Add(4)
	//h.Add(5)
	//h.Add(6)

	h = oddEvenList(h)
	ShowList(h)
}
