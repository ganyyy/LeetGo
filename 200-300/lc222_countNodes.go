//go:build ignore

package main

import . "leetgo/data"

func countNodes(root *TreeNode) int {
	// mark
	if root == nil {
		return 0
	}
	if ld, rd := getDepth(root.Left), getDepth(root.Right); ld == rd {
		// 缺口在右半部分, 左边是一个完美二叉树
		return countNodes(root.Right) + 1<<ld
	} else {
		// 缺口在左半部分, 右半部分是一颗完美二叉树
		return countNodes(root.Left) + 1<<rd
	}
}

// 根据左子树的大小来确定树的高度
func getDepth(root *TreeNode) int {
	var res int
	for root != nil {
		res++
		root = root.Left
	}
	return res
}

func countNodesAll(root *TreeNode) int {
	// 全遍历一遍是肯定可行的
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func main() {

}
