//go:build ignore
// +build ignore

package main

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	// DFS啊
	if root == nil {
		return nil
	}
	var helper func(root *Node) (tail *Node)

	helper = func(root *Node) (tail *Node) {
		if root == nil {
			return
		}
		var pre = root
		for cur := root; cur != nil; cur = cur.Next {
			if cur.Child != nil {
				var childTail = helper(cur.Child)
				childTail.Next = cur.Next
				if cur.Next != nil {
					cur.Next.Prev = childTail
				}
				cur.Next = cur.Child
				cur.Child.Prev = cur
				cur.Child = nil
				cur = childTail
			}
			pre = cur
		}
		return pre
	}

	helper(root)
	return root
}
