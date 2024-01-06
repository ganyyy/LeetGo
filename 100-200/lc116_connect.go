package main

/**
 * Definition for a Node.
 *
 *
 *
 *
 *
 *
 */

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

func connect116(root *Node116) *Node116 {
	// 层次遍历可解, 不过占用空间不是常数级别的
	// 这里采用递归的思路进行处理
	if root == nil {
		return nil
	}
	// 获取next的下一个
	var next *Node116
	if root.Next != nil {
		if root.Next.Left != nil {
			next = root.Next.Left
		} else {
			next = root.Next.Right
		}
	}
	// 赋值
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
			root.Right.Next = next
		} else {
			root.Left.Next = next
		}
	}
	// 先右边, 在左边
	connect116(root.Right)
	connect116(root.Left)
	return root
}

func connect3(root *Node116) *Node116 {
	// 这是完美二叉树, 可以简化判断条件
	if root == nil || root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	connect3(root.Left)
	connect3(root.Right)
	return root
}

func main() {

}
