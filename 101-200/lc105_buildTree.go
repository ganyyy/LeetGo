package main

import . "leetgo/data"

func buildTree(preorder []int, inorder []int) *TreeNode {
	if 0 == len(preorder) || 0 == len(inorder) {
		return nil
	}

	var root = &TreeNode{Val: preorder[0]}
	if 1 == len(preorder) {
		return root
	}

	// 找到左右分开的位置

	// 找中序的分割点
	var i int
	m := make(map[int]struct{})
	for r := preorder[0]; i < len(inorder); i++ {
		if inorder[i] == r {
			break
		} else {
			m[inorder[i]] = struct{}{}
		}
	}
	// 找前序的分割点
	var j int
	for j = 1; j < len(preorder); j++ {
		if _, ok := m[preorder[j]]; !ok {
			break
		}
	}

	if j != 1 && j <= len(preorder) && i != 0 && i <= len(inorder) {
		root.Left = buildTree(preorder[1:j], inorder[:i])
	}
	if j < len(preorder) && i+1 < len(inorder) {
		root.Right = buildTree(preorder[j:], inorder[i+1:])
	}

	return root
}

func buildTreeSimple(preorder []int, inorder []int) *TreeNode {
	if 0 == len(preorder) || 0 == len(inorder) {
		return nil
	}

	var root = &TreeNode{Val: preorder[0]}
	if 1 == len(preorder) {
		return root
	}

	// 找到左右分开的位置
	for i := 0; i < len(preorder); i++ {
		//  这就是分割点
		if preorder[0] == inorder[i] {
			root.Left = buildTree(preorder[1:i+1], inorder[:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
			break
		}
	}
	return root
}

func main() {
	buildTree([]int{1, 2}, []int{2, 1})
}
