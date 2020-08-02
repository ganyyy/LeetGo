package main

import (
	. "leetgo/data"
)

func flatten(root *TreeNode) {
	if nil == root {
		return
	}
	// 层次遍历解开即可
	var stack []*TreeNode
	var res []*TreeNode
	cur := root
	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = cur.Right
		} else {
			stack = append(stack, cur)
			res = append(res, cur)
			cur = cur.Left
		}
	}
	if len(res) == 0 {
		return
	}
	pre := res[0]
	for i := 1; i < len(res); i++ {
		pre.Left = nil
		pre.Right = res[i]
		pre = pre.Right
	}
}

func flatten2(root *TreeNode) {
	curr := root
	for curr != nil {
		// 如果当前节点的左子节点不为空, 就找到右子节点的前驱节点
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			// 前驱节点找到后, 将 当前节点的右节点赋给 前驱节点的右节点
			predecessor.Right = curr.Right
			// 当前节点的左节点清空, 右键点指向左节点
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}

func main() {
	t := &TreeNode{Val: 1}
	t.Left = &TreeNode{Val: 2}
	t.Right = &TreeNode{Val: 5}
	t.Left.Left = &TreeNode{Val: 3}
	t.Left.Right = &TreeNode{Val: 4}
	t.Right.Left = &TreeNode{Val: 6}

	flatten(t)
}
