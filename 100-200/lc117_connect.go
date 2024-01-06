//go:build ignore

package main

import (
	"fmt"
	. "leetgo/data"
)

func connect(root *TNode) *TNode {
	if root == nil {
		return root
	}

	queue := []*TNode{root}

	for len(queue) != 0 {
		ln := len(queue)
		fmt.Println(ln)
		pre := queue[0]
		for i := 0; i < ln; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i == 0 {
				continue
			}
			pre.Next = node
			pre = node
		}
		queue = queue[ln:]
	}
	return root
}

func connect2(root *TNode) *TNode {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}
	if root.Left != nil && root.Right == nil {
		root.Left.Next = getNext(root.Next)
	}
	if root.Right != nil {
		root.Right.Next = getNext(root.Next)
	}
	// 需要先对右边进行递归
	// 因为左边依赖右边的结果
	connect2(root.Right)
	connect2(root.Left)
	return root
}

// 根据父节点的next获取子节点的next
func getNext(root *TNode) *TNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		return root.Left
	}
	if root.Right != nil {
		return root.Right
	}
	if root.Next != nil {
		return getNext(root.Next)
	}
	return nil
}
