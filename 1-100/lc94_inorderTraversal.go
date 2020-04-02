package main

import (
	"fmt"
	. "leetgo/data"
)

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		if cur == nil {
			cur = stack[len(stack)-1]
			res = append(res, cur.Val)
			stack = stack[:len(stack)-1]
			cur = cur.Right
		} else {
			stack = append(stack, cur)
			cur = cur.Left
		}
	}
	//
	return res
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	fmt.Println(inorderTraversal(root))
}
