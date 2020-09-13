package main

import (
	"fmt"
	. "leetgo/data"
)

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if nil == root {
		return res
	}
	var stack []*TreeNode
	for len(stack) != 0 || root != nil {
		if root != nil {
			// 将当前节点入栈
			stack = append(stack, root)
			// 指向自己的左孩子
			root = root.Left
		} else {
			// 栈顶出栈
			ln := len(stack)
			last := stack[ln-1]
			stack = stack[:ln-1]
			// 将栈顶值输入到输出队列
			// 可以这么理解: 先找到最后一个左孩子
			// 如果左孩子没有右子节点, 那么自然会回溯到左孩子的父节点, 满足 左->中->右的输出顺序
			// 如果左孩子存在右子节点, 那么顺序就是 中->右的输出顺序
			res = append(res, last.Val)
			// 指向栈顶元素的右孩子
			root = last.Right
		}
	}
	return res
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	fmt.Println(inorderTraversal(root))
}
