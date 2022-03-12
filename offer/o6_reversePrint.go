//go:build ignore
// +build ignore

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var tmp []int
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	for l, r := 0, len(tmp)-1; l < r; l, r = l+1, r-1 {
		tmp[l], tmp[r] = tmp[r], tmp[l]
	}
	return tmp
}
