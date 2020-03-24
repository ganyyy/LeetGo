package main

import . "leetgo/data"

func rob(root *TreeNode) int {
	ret := robInternal(root)
	return max(ret[0], ret[1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robInternal(root *TreeNode) (ret [2]int) {
	// 返回两个结果:
	// ret[0] = 不选取本节点
	// ret[1] = 选取本节点
	if root == nil {
		return
	}

	left := robInternal(root.Left)
	right := robInternal(root.Right)
	ret[0] = max(left[0], left[1]) + max(right[0], right[1])
	ret[1] = left[0] + right[0] + root.Val

	return
}

func main() {

}
