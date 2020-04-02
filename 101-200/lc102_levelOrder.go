package main

import . "leetgo/data"

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	res = append(res, []int{root.Val})
	help(root, 1, &res)
	return res
}

func help(node *TreeNode, lv int, res *[][]int) {
	if node == nil {
		return
	}

	ret := make([]int, 0, 2)
	if node.Left != nil {
		ret = append(ret, node.Left.Val)
	}
	if node.Right != nil {
		ret = append(ret, node.Right.Val)
	}

	if lv >= len(*res) {
		*res = append(*res, ret)
	} else {
		(*res)[lv] = append((*res)[lv], ret...)
	}

	help(node.Left, lv+1, res)
	help(node.Right, lv+1, res)
}

func levelOrder2(root *TreeNode) [][]int {
	// 非递归调用
	if root == nil {
		return [][]int{}
	}
	// 每次都把当前队列遍历为空即可
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	var res [][]int
	for len(stack) != 0 {
		var ret []int
		for cur, ln := 0, len(stack); cur < ln; cur++ {
			front := stack[0]
			stack = stack[1:]
			ret = append(ret, front.Val)
			if front.Left != nil {
				stack = append(stack, front.Left)
			}
			if front.Right != nil {
				stack = append(stack, front.Right)
			}
		}
		res = append(res, ret)
	}
	return res
}

func main() {
	var ret []int
	ret = append(ret, 10)
}
