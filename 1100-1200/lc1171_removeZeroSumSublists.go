//go:build ignore

package main

import . "leetgo/data"

func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	seen := map[int]*ListNode{}
	prefix := 0
	for node := dummy; node != nil; node = node.Next {
		prefix += node.Val
		seen[prefix] = node
	}
	// [1,2,3,-3,1]
	// [1,3,6, 3,4]
	// [1,2,1], [6,3]被跳过
	prefix = 0
	for node := dummy; node != nil; node = node.Next {
		// 前缀和相同, 就意味着中间存在一段区间相加为0

		prefix += node.Val
		// 这里一定是不为nil的, 因为如果后续不存在, 那就即使当前节点本身, 否则就指向了下一个区间和为0的节点
		node.Next = seen[prefix].Next
	}
	return dummy.Next
}
