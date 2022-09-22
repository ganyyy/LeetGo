//go:build ignore

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&ListNode{}, 0}
}

func (l *MyLinkedList) Get(index int) int {
	if index < 0 || index >= l.size {
		return -1
	}
	return l.preNode(index + 1).Val
}

func (l *MyLinkedList) AddAtHead(val int) {
	l.AddAtIndex(0, val)
}

func (l *MyLinkedList) AddAtTail(val int) {
	l.AddAtIndex(l.size, val)
}

func (l *MyLinkedList) preNode(index int) *ListNode {
	pre := l.head
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	return pre
}

func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index > l.size {
		return
	}
	index = max(index, 0)
	l.size++
	preNode := l.preNode(index)
	toAdd := &ListNode{val, preNode.Next}
	preNode.Next = toAdd
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= l.size {
		return
	}
	l.size--
	preNode := l.preNode(index)
	preNode.Next = preNode.Next.Next
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
