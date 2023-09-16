package main

import . "leetgo/data"

func deleteDuplicates3(head *ListNode) *ListNode {
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

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// var h = &ListNode{Val: head.Val - 1, Next: head}
	var h = &ListNode{Next: head}
	pre, head := h, head.Next
	var remove bool
	// 严格意义上来讲, 和之前额没啥区别. 不存在更差的可能
	for ; head != nil; head = head.Next {
		if pre.Next.Val == head.Val {
			remove = true
			continue
		}
		if remove {
			pre.Next = head
			remove = false
		} else {
			pre = pre.Next
		}
	}
	if remove {
		pre.Next = nil
	}
	return h.Next
}

func main() {
	var head = &ListNode{
		Val:  1,
		Next: nil,
	}
	head.Add(1)
	head.Add(2)
	// head.Add(2)
	// head.Add(3)

	head.Add(3)
	head.Add(4)
	// head.Add(4)
	// head.Add(5)
	head.Add(5)
	ShowList(head)
	ShowList(deleteDuplicates2(head))
}
