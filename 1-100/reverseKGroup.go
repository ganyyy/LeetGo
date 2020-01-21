package main

import . "leetgo/data"

func reverseKGroup(head *ListNode, k int) *ListNode {
	n := &ListNode{
		Val:  -1,
		Next: head,
	}
	gHead, now, count := n, head, 1
	for nil != now {
		if count >= k {
			// 第一次逆转两个位置
			count -= 1
			// 逆转
			// 这里有单链表的逆序
			// 核心想法是:
			// t 为s节点的下一个
			// tn 为t的下一个
			// t.n = tn.n 将t的下一个指向tn的下一个
			// tn.n = s.n tn的下一个指向s节点的下一个
			// s.n = tn s节点的下一个指向tn
			// 此时t不需要重新赋值, 因为t的下一个就已经是新节点了
			// 以后的操作就相当一向头部插入新节点, 尾部一直向下指
			t := gHead.Next
			for count > 0 {
				tn := t.Next
				t.Next = tn.Next
				tn.Next = gHead.Next
				gHead.Next = tn
				count--
			}
			// 重置一下
			count = 0
			// 下一组的头
			gHead, now = t, t
		}
		now = now.Next
		count++
	}

	return n.Next
}

func main() {
	h := &ListNode{
		Val:  0,
		Next: nil,
	}

	for i := 1; i <= 6; i++ {
		h.Add(i)
	}

	ShowList(h)

	h = reverseKGroup(h, 2)

	ShowList(h)
}
