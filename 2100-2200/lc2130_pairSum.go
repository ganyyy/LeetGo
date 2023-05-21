//go:build ignore

package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// 快慢指针找中点

	if head == nil {
		return 0
	}
	var slow, fast = head, head.Next
	// fmt.Println("start mid")
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// fmt.Println("start reverse")
	mid := slow.Next
	slow.Next = nil
	// 反转 [head, slow]
	nxt := head.Next
	head.Next = nil
	for nxt != nil {
		nn := nxt.Next
		nxt.Next = head
		head = nxt
		nxt = nn
	}

	// fmt.Println("start ret")
	// 求最大值
	var ret = math.MinInt32

	for mid != nil && head != nil {
		sum := mid.Val + head.Val
		if sum > ret {
			ret = sum
		}
		mid = mid.Next
		head = head.Next
	}

	return ret
}
