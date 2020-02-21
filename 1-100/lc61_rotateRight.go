package main

import . "leetgo/data"

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 先确定数量, 在确定最终偏移量, 时间复杂度为O(n)
	tmp, pre, count := head, (*ListNode)(nil), 0
	for tmp != nil {
		pre = tmp
		tmp = tmp.Next
		count++
	}
	// 形成环
	pre.Next = head
	// 确定实际偏移量
	k %= count
	k = count - k
	tmp = head
	for i := 0; i < k; i++ {
		pre = tmp
		tmp = tmp.Next
	}
	// 解开环
	pre.Next = nil
	return tmp
}

func main() {
	head := &ListNode{
		Val: 1,
	}
	head = rotateRight(head, 4)
	ShowList(head)
}
