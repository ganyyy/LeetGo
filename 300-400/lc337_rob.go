//go:build ignore

package main

import . "leetgo/data"

func rob(root *TreeNode) int {
	// mark
	ret := robHelper(root)
	return max(ret[0], ret[1])
}

// 两个值, 偷当前房子和不偷当前房子
func robHelper(root *TreeNode) (ret [2]int) {
	if nil == root {
		return
	}
	// 获取两个孩子的最大价值
	nextLeft := robHelper(root.Left)
	nextRight := robHelper(root.Right)
	// 如果选择取当前节点, 那么两个孩子节点就不能取
	ret[0] = root.Val + nextLeft[1] + nextRight[1]
	// 如果不选择当前节点, 就取 两个孩子中的 取和不取 之间的最大值和
	ret[1] = max(nextLeft[0], nextLeft[1]) + max(nextRight[0], nextRight[1])
	return
}

func main() {

}
