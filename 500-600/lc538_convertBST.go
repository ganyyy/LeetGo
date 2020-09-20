package main

import . "leetgo/data"

var num int

func convertBST(root *TreeNode) *TreeNode {
	// 中序遍历解决问题
	// 二叉排序树 的中序遍历为从小到大, 反过来就是从大到小
	// 累加即可
	if root == nil {
		return root
	}
	convertBST(root.Right)
	root.Val += num
	num = root.Val
	convertBST(root.Left)
	return root
}
