package main

import "sort"

import . "leetgo/data"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortListSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 不可取, 但是肯定还是要缓存所有的节点, 不然可能会出现 n^2 的时间复杂度
	// 搞一个数组, 存放所有的节点, 然后排序, 然后在组合起来?
	var nodes = make([]*ListNode, 0, 100)

	// n
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	// nlog(n)
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	// n
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil
	return nodes[0]
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1. 保留头部
	// 2. 保留当前排序所处链表末尾
	var tail, cur = head, head.Next
	// 构建一个指向头节点的节点
	var pre = &ListNode{Next: head}
	for cur != nil {
		var now = cur
		cur = cur.Next
		now.Next = nil
		// 如果当前节点大于已排序列表的末尾节点, 直接插入到后边
		if now.Val > tail.Val {
			tail.Next = now
			tail = tail.Next
			continue
		}
		var p = pre
		// 从头开始查找当前节点应处的位置
		for head.Val < now.Val {
			p = head
			head = head.Next
		}
		now.Next = p.Next
		p.Next = now
		head = pre.Next
	}
	// 断掉结尾的下一个节点
	tail.Next = nil
	return pre.Next
}

func main() {
	var l = &ListNode{Val: 4}
	l.Add(2).Add(1).Add(3)
	ShowList(insertionSortList(l))
}
