//go:build ignore
// +build ignore

package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	// 递归式写法

	var sb strings.Builder

	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		sb.WriteString(strconv.Itoa(root.Val))
		if root.Left != nil {
			sb.WriteString("(" + tree2str(root.Left) + ")")
		}
		if root.Right != nil {
			if root.Left == nil {
				sb.WriteString("()")
			}
			sb.WriteString("(" + tree2str(root.Right) + ")")
		}
		return
	}

	dfs(root)

	return sb.String()
}
