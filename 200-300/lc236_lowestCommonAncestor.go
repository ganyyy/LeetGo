package main

import . "leetgo/data"

// 解法1
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 叶节点

	if nil == root {
		return root
	}
	// 找到了其中一个节点, 直接返回即可
	if p == root || q == root {
		return root
	}

	// 分别找左右两边
	left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
	// 左右节点都在, 说明这是个公共节点
	if nil != left && nil != right {
		return root
	}
	// 返回左右节点不为空的即可
	if nil != left {
		return left
	}
	// 不管右节点存不存在都无所谓了
	return right
}

// 解法2
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {

	var ans *TreeNode
	var search func(r *TreeNode) bool

	search = func(r *TreeNode) bool {
		if r == nil {
			return false
		}

		var self, left, right int

		if r == p || r == q {
			self = 1
		}

		if search(r.Left) {
			left = 1
		}

		if search(r.Right) {
			right = 1
		}

		// 如果满足任意一种组合, 即 left+right, left/right + self
		// 既可以视为满足要求, 此时就是正解
		if (left + self + right) >= 2 {
			ans = r
		}

		return (left + self + right) > 0
	}

	search(root)
	return ans
}

func main() {

}
