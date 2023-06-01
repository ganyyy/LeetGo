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
	// 递归的经典应用
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
		// 所谓的后继节点, 就是当前节点右子树的最左边的节点
		successor := root.Right
		for successor.Left != nil {
			successor = successor.Left
		}
		// 移除root后, 将successor替换到root的位置上
		// 首先, 因为successor来自于Right, 那么就需要从Right中删除
		successor.Right = deleteNode(root.Right, successor.Val)
		// 然后, 继承root.Left
		successor.Left = root.Left
		return successor
	}
	return root
}

func deleteNode3(root *TreeNode, key int) *TreeNode {
	// cur: 当前节点(Val == key)
	// curParent: 父节点(Left/Right = cur)
	var cur, curParent *TreeNode = root, nil
	for cur != nil && cur.Val != key {
		curParent = cur
		if cur.Val > key {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	if cur == nil {
		// 找不到
		return root
	}
	if cur.Left == nil && cur.Right == nil {
		// 这是个叶子节点
		cur = nil
	} else if cur.Right == nil {
		// 右子树为空
		cur = cur.Left
	} else if cur.Left == nil {
		// 左子树为空
		cur = cur.Right
	} else {
		// 找到后继节点, 替换当前节点的位置
		successor, successorParent := cur.Right, cur
		for successor.Left != nil {
			// 后继节点就是右子树的最小值, 一直向左找到头即可
			successorParent = successor
			successor = successor.Left
		}
		if successorParent.Val == cur.Val {
			// 特殊情况: cur.Right 就是一个叶子节点
			successorParent.Right = successor.Right
		} else {
			// 将后继节点右子树接到前驱节点的左子树上
			// 后继节点是肯定不会存在左子树的, 因为这已经左到头了
			successorParent.Left = successor.Right
		}
		// 将后继节点替换到当前节点
		successor.Right = cur.Right
		successor.Left = cur.Left
		cur = successor
	}
	if curParent == nil {
		// 如果要删除的就是根节点呢?
		return cur
	}
	// 替换被删除节点的父节点的指向节点
	if curParent.Left != nil && curParent.Left.Val == key {
		curParent.Left = cur
	} else {
		curParent.Right = cur
	}
	return root
}

func main() {
	t := &TreeNode{}
	t.Val = 1
	t.Right = &TreeNode{Val: 2}

	deleteNode(t, 2)
}
