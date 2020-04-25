package main

import "math"
import . "leetgo/data"

func isValidBST(root *TreeNode) bool {
	if nil == root {
		return true
	}
	return checkValid(root, math.MinInt64, math.MaxInt64)
}

func checkValid(root *TreeNode, min, max int) bool {
	if nil == root {
		return true
	}
	if min >= root.Val || max <= root.Val {
		return false
	}
	return checkValid(root.Left, min, root.Val) && checkValid(root.Right, root.Val, max)
}

func main() {

}
