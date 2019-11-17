package main

import "fmt"

func (l *ListNode) Add(val int) {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	s := l
	for s.Next != nil {
		s = s.Next
	}
	s.Next = node
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var left = head
	if nil == head {
		return nil
	}
	count := 0
	now := head
	for now != nil {
		if count > n {
			left = left.Next
		}
		count++
		now = now.Next
	}
	if count == n {
		return head.Next
	} else {
		if count > n {
			left.Next = left.Next.Next
		}
		return head
	}
}

func AddNode(before *ListNode, val int) *ListNode {
	next := &ListNode{
		Val:  val,
		Next: nil,
	}
	before.Next = next
	return next
}

func ShowList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	next := head
	for _i := 2; _i <= 5; _i++ {
		next = AddNode(next, _i)
	}

	ShowList(head)

	fmt.Println()

	head = removeNthFromEnd(head, 6)

	ShowList(head)
}
