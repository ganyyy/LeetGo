//go:build ignore

package main

import . "leetgo/data"

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	haveSufficient := checkSufficientLeaf(root, 0, limit)
	if haveSufficient {
		return root
	} else {
		return nil
	}
}

func checkSufficientLeaf(node *TreeNode, sum int, limit int) bool {
	if node == nil {
		return false
	}
	cur := node.Val + sum
	if node.Left == nil && node.Right == nil {
		// 先找到叶子节点再说
		return cur >= limit
	}
	haveSufficientLeft := checkSufficientLeaf(node.Left, cur, limit)
	haveSufficientRight := checkSufficientLeaf(node.Right, cur, limit)
	if !haveSufficientLeft {
		node.Left = nil
	}
	if !haveSufficientRight {
		node.Right = nil
	}
	return haveSufficientLeft || haveSufficientRight
}
