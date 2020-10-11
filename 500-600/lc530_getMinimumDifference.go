package main

import "math"

import . "leetgo/data"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	// 中序走一波?
	// var stack = []*TreeNode{}
	var res int = math.MaxInt32
	var pre = -1
	// for len(stack) != 0 || root != nil {
	//     if root == nil {
	//         top := stack[len(stack)-1]
	//         stack = stack[:len(stack)-1]
	//         if pre != -1 {
	//             res = min(res, abs(top.Val, pre))
	//         }
	//         pre = top.Val
	//         root = top.Right
	//     } else {
	//         stack = append(stack, root)
	//         root = root.Left
	//     }
	// }

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != -1 {
			res = min(res, abs(root.Val, pre))
		}
		pre = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {

}
