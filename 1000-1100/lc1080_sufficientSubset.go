//go:build ignore

package main

import . "leetgo/data"

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	// 注意: 从根节点到叶子节点的路径和必须大于等于 limit, 所以如果某个节点的子树和小于 limit, 那么这个节点就要被删除
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
		// 叶子节点, 判断是否满足条件
		return cur >= limit
	}
	// 找到左右子树是否满足条件
	haveSufficientLeft := checkSufficientLeaf(node.Left, cur, limit)
	haveSufficientRight := checkSufficientLeaf(node.Right, cur, limit)
	if !haveSufficientLeft {
		// 如果左子树不满足条件, 那么左子树就要被删除
		node.Left = nil
	}
	if !haveSufficientRight {
		// 如果右子树不满足条件, 那么右子树就要被删除
		node.Right = nil
	}
	// 如果左右子树都不满足条件, 那么当前节点就要被删除
	// 如果左右子树至少有一个满足条件, 那么当前节点就要保留
	return haveSufficientLeft || haveSufficientRight
}
