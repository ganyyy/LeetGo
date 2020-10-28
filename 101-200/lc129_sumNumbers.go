package main

import . "leetgo/data"

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int

	var helper func(root *TreeNode, pre int)

	helper = func(root *TreeNode, pre int) {
		// 理论上这种情况应该不会出现
		// if root == nil {
		//     res += pre
		//     return
		// }
		// 叶子节点
		if root.Left == nil && root.Right == nil {
			res += pre*10 + root.Val
			return
		}
		// 看左子节点
		if root.Left != nil {
			helper(root.Left, pre*10+root.Val)
		}
		// 看右子节点
		if root.Right != nil {
			helper(root.Right, pre*10+root.Val)
		}
	}

	helper(root, 0)

	return res
}
