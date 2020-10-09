package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 设链表头->环入口的长度为a, 设环长度为b
	// 快指针走两步, 慢指针走一步, 设相遇在 距离a x的位置,
	// 此时 快指针走了 a + mb + x 步(m >= 1)
	// 满指针走了 a + x 步(a+x<a+b)
	// 又∵ 快指针走过的路程是满指针的2倍, 所以有
	// 2(a+x) = a+mb+x => a = mb - x
	// 想要求得 a 的位置, 此时慢满指针指向head, 令慢指针和快指针一起走, 当二者相遇时, 交点就是 a
	// 理由如下:
	// 假设 二者走了 y 步发生了相遇, 那么 快指针一共走了 a+mb+x+y 步, 带入 a = mb-x 可得 2mb+y
	// 很明显 mb是环的大小. 当且仅当 y = a时才符合快指针走的总路程.

	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	l := &ListNode{Val: 3}
	p := l.Add(2)
	l.Add(0)
	q := l.Add(-4)
	q.Next = p
	detectCycle(l)
}
