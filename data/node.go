package data

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Add(val int) *ListNode {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	s := l
	for s.Next != nil {
		s = s.Next
	}
	s.Next = node
	return node
}

func ShowList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d", head.Val)
		head = head.Next
		if nil != head {
			fmt.Print("->")
		}
	}
	fmt.Println()
}
