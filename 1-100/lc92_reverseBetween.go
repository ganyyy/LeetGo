package main

import . "leetgo/data"

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m >= n || nil == head {
		return head
	}
	var count = 1
	var pre = &ListNode{Next: head}
	now := head
	if m != 1 {
		for count < m-1 {
			now = now.Next
			count++
		}
		pre = now
		now = now.Next
		count++
	}
	for count < n && now != nil {
		// 核心是这一块, 先取下一个节点
		nn := now.Next
		// 将当前位置指向下一个节点的下一个节点
		now.Next = nn.Next
		// 下一个节点指向前置节点的下一个
		nn.Next = pre.Next
		// 根节点下一个指向nn
		pre.Next = nn
		count++
	}
	if m == 1 {
		return pre.Next
	} else {
		return head
	}
}

func main() {
	head := &ListNode{Val: 1}
	head.Add(2)
	head.Add(3)
	head.Add(4)
	head.Add(5)
	ShowList(reverseBetween(head, 1, 5))
}
