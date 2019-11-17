package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l2 {
		return l1
	}
	if nil == l1 {
		return l2
	}
	// 保证l1是小的开头的那一组
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	//  尝试以 l1做最后的返回链表
	head := l1
	// l1的最后节点的前一个节点
	pre := l1
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			// 如果发现当前l1比l2要大
			// l1前节点的next指向l2
			pre.Next = l2
			// l2指向l2的下一个
			l2 = l2.Next
			// 复原l1
			pre.Next.Next = l1
			// 前置节点位置更新
			pre = pre.Next
		} else {
			// 如果当前l1比l2小, 那么正常轮询即可
			pre = l1
			l1 = l1.Next
		}
	}
	if l2 != nil {
		if l1 != nil {
			l1.Next = l2
		} else {
			pre.Next = l2
		}
	}
	return head
}

func ShowList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l1.Add(10)
	l1.Add(20)
	ShowList(l1)
	fmt.Println()

	l2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2.Add(3)
	l2.Add(4)
	l2.Add(6)
	l2.Add(10)
	ShowList(l2)

	fmt.Println()
	res := mergeTwoLists(l1, l2)
	ShowList(res)
}
