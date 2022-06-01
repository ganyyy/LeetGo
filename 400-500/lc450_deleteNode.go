//go:build ignore

package main

import . "leetgo/data"

func deleteNode(root *TreeNode, key int) *TreeNode {
	if nil == root {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	} else {
		// 此时就是删除当前节点

		// 右子树为空, 返回左子树
		if nil == root.Right {
			return root.Left
		}
		// 左子树为空, 返回右子树
		if nil == root.Left {
			return root.Right
		}

		// 两边子树都不为空

		minNode := root.Right
		// 找后继, 即右子树的最左边
		for nil != minNode.Left {
			minNode = minNode.Left
		}
		root.Val = minNode.Val
		root.Right = delMinNode(root.Right)
		return root
	}
}

func delMinNode(root *TreeNode) *TreeNode {
	// 如果左子树为空, 直接删除接上右子树即可
	if nil == root.Left {
		return root.Right
	}
	// 递归的将找没有左节点的左子树
	root.Left = delMinNode(root.Left)
	return root
}

func deleteNode2(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	case root.Left == nil || root.Right == nil:
		// 找到节点了, 此时左右子节点有一个是空的, 返回非空的那一个
		if root.Left != nil {
			return root.Left
		}
		return root.Right
	default:
		// 左右子节点都不是空的, 这里寻找的是后继节点
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}

func main() {
	t := &TreeNode{}
	t.Val = 1
	t.Right = &TreeNode{Val: 2}

	deleteNode(t, 2)
}
