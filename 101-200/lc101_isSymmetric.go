package main

import (
	. "leetgo/data"
)

func isSymmetric(root *TreeNode) bool {
	// 注意空树的情况
	if nil == root {
		return true
	}
	// 层次遍历, 然后看是不是对称的
	stack := []*TreeNode{
		root,
	}
	for t := len(stack); 0 != t; t = len(stack) {
		left, right := 0, t-1
		// 再看看是不是对称的
		for left < right {
			l, r := stack[left], stack[right]
			if (l != nil && r != nil && l.Val == r.Val) || (l == nil && r == nil) {
				left++
				right--
			} else {
				return false
			}
		}
		// 层次遍历下一层
		for i := 0; i < t; i++ {
			top := stack[0]
			stack = stack[1:]
			if nil != top {
				stack = append(stack, top.Left)
				stack = append(stack, top.Right)
			}
		}
	}
	return true
}

// 递归版
func isSymmetric2(root *TreeNode) bool {
	var fn func(left *TreeNode, right *TreeNode) bool

	// 核心点是找到子问题, 然后对其进行求解
	fn = func(left *TreeNode, right *TreeNode) bool {
		// 都为空, 说明对称
		if nil == left && nil == right {
			return true
		}
		// 一个不为空, 说明不对称
		if nil == left || nil == right {
			return false
		}
		// 如果相等, 进行下一轮的比较
		if left.Val == right.Val {
			// 按照对称关系来
			return fn(left.Left, right.Right) && fn(left.Right, right.Left)
		} else {
			// 不相等直接返回错误
			return false
		}
	}

	return fn(root, root)
}

func main() {

}
