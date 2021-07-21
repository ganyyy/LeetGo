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

	curr := head
	for curr != nil {
		// 复制原始节点, 将新节点插入到原始节点之间
		next := curr.Next
		var clone *Node
		clone = new(Node)
		clone.Val = curr.Val
		clone.Next = next
		curr.Next = clone
		curr = next
	}

	curr = head
	for curr != nil {
		newCurr := curr.Next
		if curr.Random != nil {
			// curr.Random.Next 指向的就是复制后的新节点, 所以可以直接赋值
			newCurr.Random = curr.Random.Next
		} else {
			// 不存在就表示随机节点为空
			newCurr.Random = nil
		}
		curr = newCurr.Next
	}

	// 将新旧节点组成的链表进行拆分
	newHead := head.Next
	curr = head
	for curr != nil {
		newCurr := curr.Next
		next := newCurr.Next
		curr.Next = next
		if next != nil {
			newCurr.Next = next.Next
		} else {
			newCurr.Next = nil
		}
		curr = next
	}
	return newHead
}
