//go:build ignore

package main

import . "leetgo/data"

func searchBST(root *TreeNode, val int) *TreeNode {
	if nil == root {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	} else if root.Val < val {
		return searchBST(root.Right, val)
	} else {
		return root
	}
}

// 非递归版
func searchBST2(root *TreeNode, val int) *TreeNode {
	for nil != root {
		if val == root.Val {
			return root
		}
		if val > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return nil
}

// 还是非递归版好点
