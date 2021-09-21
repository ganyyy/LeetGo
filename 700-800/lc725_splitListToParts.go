package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	var cnt int
	for cur := head; cur != nil; cur = cur.Next {
		cnt++
	}
	var num = cnt / k
	var remain = cnt % k

	var ret = make([]*ListNode, k)

	var cur = head
	for i := 0; i < k && cur != nil; i++ {
		cnt = num
		if remain > 0 {
			remain--
			cnt++
		}
		var tmp = cur
		var pre *ListNode
		for j := 0; j < cnt && cur != nil; j++ {
			pre = cur
			cur = cur.Next
		}
		if pre != nil {
			pre.Next = nil
		}
		ret[i] = tmp
	}
	return ret
}
