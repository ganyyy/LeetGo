package main

import . "leetgo/data"

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var helper func(root *TreeNode, val int)
	var res [][]int
	var tmp []int
	helper = func(root *TreeNode, val int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			// 叶子节点
			if val < root.Val {
				return
			}
			if val == root.Val {
				t := make([]int, len(tmp)+1)
				copy(t, tmp)
				t[len(tmp)] = val
				res = append(res, t)
			}
		} else {
			tmp = append(tmp, root.Val)
			helper(root.Left, val-root.Val)
			helper(root.Right, val-root.Val)
			tmp = tmp[:len(tmp)-1]
		}

	}
	helper(root, sum)
	return res
}
