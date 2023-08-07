package main

/**
 * Definition for a Node.
 *
 *
 *
 *
 *
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomListBad(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 消耗临时空间.
	var arr []*Node
	var m = make(map[*Node]int)
	var idx int
	// 记录基础顺序和关系映射
	for cur := head; cur != nil; cur = cur.Next {
		m[cur] = idx
		idx++
		arr = append(arr, &Node{Val: cur.Val})
		if len(arr) > 1 {
			arr[len(arr)-2].Next = arr[len(arr)-1]
		}
	}

	// 构建random节点
	idx = 0
	for cur := head; cur != nil; cur = cur.Next {
		if r, ok := m[cur.Random]; ok {
			arr[idx].Random = arr[r]
		}
		idx++
	}

	return arr[0]
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// mark

	// 使用插入的方式, 将复制节点串联到原始节点后面
	var cur = head
	for cur != nil {
		var next = cur.Next
		var newNode = &Node{Val: cur.Val, Next: next}
		cur.Next = newNode
		cur = next
	}

	// 更新Random节点
	cur = head
	for cur != nil {
		var newCurr = cur.Next
		if cur.Random != nil {
			// 将Random节点更新成复制的节点
			newCurr.Random = cur.Random.Next
		}
		// 下一个正常节点
		cur = newCurr.Next
	}

	var newHead = head.Next
	cur = head
	// 拆分原始节点和复制节点
	for cur != nil {
		var newCur = cur.Next
		var oldNext = newCur.Next
		cur.Next = oldNext
		if oldNext != nil {
			newCur.Next = oldNext.Next
		} else {
			newCur.Next = nil
		}
		cur = cur.Next
	}
	return newHead
}
