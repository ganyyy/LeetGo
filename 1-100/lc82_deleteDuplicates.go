package main

import . "leetgo/data"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cur := &ListNode{}
	cur.Next = head
	ret := cur
	for head != nil && head.Next != nil {
		if head.Next.Val == head.Val {
			label := head.Val
			// 去掉所有相同的节点
			for head != nil && head.Val == label {
				head = head.Next
				cur.Next = head
			}
		} else {
			// 到这里一定是不相同的, 直接加上就好了
			cur = cur.Next
			head = head.Next
		}
	}
	return ret.Next
}

func main() {
	var head = &ListNode{
		Val:  1,
		Next: nil,
	}
	head.Add(1)
	head.Add(1)

	head.Add(2)
	head.Add(3)
	head.Add(3)
	head.Add(3)
	ShowList(head)
	ShowList(deleteDuplicates(head))
}
