package main

import . "leetgo/data"

func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	// 利用二叉搜索树的特性
	// 如果p, q 的值都小于 root, 说明 p, q 在左子树, 反之在右子树
	// 如果 p -root * q - root < 0 说明 一个在左, 一个在右, 要找的就是这个节点

	if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
		return root
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor3(root.Right, p, q)
	} else {
		return lowestCommonAncestor3(root.Left, p, q)
	}
}
