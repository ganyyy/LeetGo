package main

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// 快慢指针, 没病走两步
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		fast = fast.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func main() {

}
