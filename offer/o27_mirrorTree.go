package main

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  mirrorTree(root.Right),
		Right: mirrorTree(root.Left),
	}
}
