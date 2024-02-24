package main

import (
	. "leetgo/data"

	"math"
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestNodesDFS(root *TreeNode, queries []int) [][]int {
	var result [][]int
	var search func(root *TreeNode, query, minVal, maxVal int)
	search = func(root *TreeNode, query, minVal, maxVal int) {
		if root == nil {
			// 找不到了, 基于最大/最小值组装结果
			var ret = make([]int, 2)
			if minVal == math.MinInt32 {
				ret[0] = -1
			} else {
				ret[0] = minVal
			}
			if maxVal == math.MaxInt32 {
				ret[1] = -1
			} else {
				ret[1] = maxVal
			}
			result = append(result, ret)
			return
		}
		if root.Val == query {
			result = append(result, []int{query, query})
			return
		} else if root.Val > query {
			// 向左
			search(root.Left, query, minVal, root.Val)
		} else {
			// 向右
			search(root.Right, query, root.Val, maxVal)
		}
	}

	for _, query := range queries {
		search(root, query, math.MinInt32, math.MaxInt32)
	}
	return result
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	var a []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		a = append(a, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	ans := make([][]int, len(queries))
	for i, q := range queries {
		mn, mx := -1, -1
		j, ok := slices.BinarySearch(a, q)
		if j < len(a) {
			mx = a[j]
		}
		if !ok { // a[j]>q, a[j-1]<q
			j--
		}
		if j >= 0 {
			mn = a[j]
		}
		ans[i] = []int{mn, mx}
	}
	return ans
}
