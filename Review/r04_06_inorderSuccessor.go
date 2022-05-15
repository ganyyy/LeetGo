package main

import . "leetgo/data"

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	// 老老实实中序迭代吧...
	var stack []*TreeNode
	var find bool
	for len(stack) != 0 || root != nil {
		// fmt.Println(root)
		if root == nil {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root.Val == p.Val {
				find = true
			} else if find {
				return root
			}
			root = root.Right
		} else {
			stack = append(stack, root)
			root = root.Left
		}
	}
	return nil
}

func inorderSuccessorGood(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val >= root.Val {
		return inorderSuccessor(root.Right, p)
	}

	// 如果找得到, 那么找到的节点就是后继节点
	// 如果找不到, 那么根节点就是后继
	var next = inorderSuccessor(root.Left, p)
	if next == nil {
		return root
	}
	return next
}
