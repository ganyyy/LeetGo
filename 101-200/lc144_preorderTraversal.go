package main

import (
	. "leetgo/data"
)

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = cur.Right
		} else {
			stack = append(stack, cur)
			res = append(res, cur.Val)
			cur = cur.Left
		}
	}
	//
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	var stack []*TreeNode

	var res []int
	var top *TreeNode
	for len(stack) != 0 || root != nil {
		if root == nil {
			top = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = top.Right
		} else {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
	}
	return res
}

func main() {

}
