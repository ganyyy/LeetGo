package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var set = make(map[*Node]*Node)

	var ret = &Node{}
	var cur = ret
	var curSrc = head
	for head != nil {
		var node = &Node{
			Val: head.Val,
		}
		set[head] = node
		cur.Next = node
		head = head.Next
		cur = node
	}

	head = curSrc
	for head != nil {
		var copyNode = set[head]
		copyNode.Random = set[head.Random]
		head = head.Next
	}
	return ret.Next
}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return head
	}
	var srcHead = head
	var copyNode *Node
	for head != nil {
		// copyNode
		copyNode = &Node{
			Val:  head.Val,
			Next: head.Next,
		}
		head.Next = copyNode
		head = copyNode.Next
	}

	head = srcHead

	for head != nil {
		// 更新random指针
		copyNode = head.Next
		if head.Random != nil {
			copyNode.Random = head.Random.Next
		}
		head = copyNode.Next
	}

	var newHead = srcHead.Next
	var cur = srcHead
	// 拆分原始节点和复制节点
	// a->a'->b->b'->null
	for cur != nil {
		// a'
		var newCur = cur.Next
		// b
		var oldNext = newCur.Next
		// a->b
		cur.Next = oldNext
		// a'->b'
		if oldNext != nil {
			newCur.Next = oldNext.Next
		} else {
			newCur.Next = nil
		}
		// b
		cur = cur.Next
	}
	return newHead
}
