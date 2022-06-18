package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 * }
 */

func insertBad(aNode *Node, x int) *Node {
	if aNode == nil {
		var n = &Node{Val: x}
		n.Next = n
		return n
	}

	var min = aNode
	// 首个节点也要看一下啊
	var find = min.Val <= x
	for cur := aNode.Next; cur != aNode; cur = cur.Next {
		// fmt.Printf("cur: %+v\n", cur)
		/*
		   1. 找到第一个小于x的节点
		   2. 找到小于x的最大节点
		*/
		if (min.Val > x && cur.Val <= x) || (cur.Val <= x && min.Val <= cur.Val) {
			// fmt.Printf("new min: %+v\n", min)
			find = true
			min = cur
		} else if !find && min.Val <= cur.Val {
			min = cur
		}
	}
	// fmt.Printf("final min: %+v\n", min)
	var n = &Node{Val: x}
	n.Next = min.Next
	min.Next = n
	return aNode
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil { // 链表为空
		aNode = &Node{Val: x}
		aNode.Next = aNode
		return aNode
	}
	for p, n := aNode, aNode.Next; ; p, n = p.Next, n.Next {
		// 断崖
		// 升序中
		// 链表等高
		if (p.Val > n.Val && (p.Val <= x || x <= n.Val)) ||
			(p.Val <= x && x <= n.Val) ||
			n == aNode {
			p.Next = &Node{Val: x, Next: p.Next}
			return aNode
		}
	}
}
