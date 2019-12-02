package main

import "fmt"
import . "leetgo/data"

func mergeKLists(lists []*ListNode) *ListNode {
	total := len(lists)
	if total == 0 {
		return nil
	}
	if total < 2 {
		return lists[0]
	}

	interval := 1

	for interval < total {
		for i := 0; i < total-interval; i += interval*2 {
			lists[i] = mergeTwoList(lists[i], lists[i+interval])
		}
		interval *= 2
	}

	return lists[0]
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	if nil == list1 {
		return list2
	}
	if nil == list2 {
		return list1
	}
	// 1始终是最小的
	if list1.Val > list2.Val {
		list1, list2 = list2, list1
	}
	head := list1
	pre := head
	list1 = list1.Next
	for nil != list1 && nil != list2 {
		if list1.Val < list2.Val {
			pre = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
			pre.Next.Next = list1
			pre = pre.Next
		}
	}

	if nil != list2 {
		if nil != list1 {
			list1.Next = list2
		} else {
			pre.Next = list2
		}
	}
	return head
}


// 想法是两两合并

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l1.Add(10)
	l1.Add(20)
	ShowList(l1)
	fmt.Println()

	l2 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2.Add(3)
	l2.Add(4)
	l2.Add(6)
	l2.Add(10)
	ShowList(l2)
	fmt.Println()

	l3 := &ListNode{
		Val:  7,
		Next: nil,
	}
	l3.Add(10)
	l3.Add(12)
	ShowList(l3)
	fmt.Println()
	res := mergeKLists([]*ListNode{l1, l2, l3})
	ShowList(res)
}
