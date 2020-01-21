package main

import "fmt"
import . "leetgo/data"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var left = head
	if nil == head {
		return nil
	}
	count := 0
	now := head
	for now != nil {
		if count > n {
			left = left.Next
		}
		count++
		now = now.Next
	}
	if count == n {
		return head.Next
	} else {
		if count > n {
			left.Next = left.Next.Next
		}
		return head
	}
}

func main() {
	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	for _i := 2; _i <= 5; _i++ {
		head.Add(_i)
	}

	ShowList(head)

	fmt.Println()

	head = removeNthFromEnd(head, 6)

	ShowList(head)
}
