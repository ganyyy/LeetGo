package main

import (
	. "leetgo/data"
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type u1 struct {
	row int // 所出现的行
	val int // 值
}

type u2 struct {
	col  int
	vals []u1
}

func verticalTraversal(root *TreeNode) [][]int {
	var m = make(map[int][]u1)

	var dfs func(root *TreeNode, col, row int)

	dfs = func(root *TreeNode, col, row int) {
		if root == nil {
			return
		}
		m[col] = append(m[col], u1{val: root.Val, row: row})
		dfs(root.Left, col-1, row+1)
		dfs(root.Right, col+1, row+1)
	}

	dfs(root, 0, 0)

	// 构建答案

	var units = make([]u2, 0, len(m))
	for k, v := range m {
		sort.Slice(v, func(i, j int) bool {
			if v[i].row != v[j].row {
				return v[i].row < v[j].row
			}
			return v[i].val < v[j].val
		})
		units = append(units, u2{
			col:  k,
			vals: v,
		})
	}
	sort.Slice(units, func(i, j int) bool {
		return units[i].col < units[j].col
	})

	var ret = make([][]int, len(units))
	for i, v := range units {
		var tmp = make([]int, len(v.vals))
		for j, v2 := range v.vals {
			tmp[j] = v2.val
		}
		ret[i] = tmp
	}
	return ret
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func verticalTraversal2(root *TreeNode) [][]int {
	const POW = 15 // 最大值1000, 10000作为基数肯定够了
	const BASE = (1 << POW) - 1

	var m = make(map[int][]int)

	var dfs func(root *TreeNode, col, row int)
	dfs = func(root *TreeNode, col, row int) {
		if root == nil {
			return
		}
		m[col] = append(m[col], row<<POW|root.Val)
		dfs(root.Left, col-1, row+1)
		dfs(root.Right, col+1, row+1)
	}

	dfs(root, 0, 0)

	type unix struct {
		col  int
		vals []int
	}

	var tmp = make([]unix, 0, len(m))
	for col, vals := range m {
		sort.Ints(vals)
		tmp = append(tmp, unix{
			col:  col,
			vals: vals,
		})
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].col < tmp[j].col
	})

	var ret = make([][]int, len(tmp))

	for i, u := range tmp {
		var t = make([]int, len(u.vals))
		for j, v := range u.vals {
			t[j] = v & BASE
		}
		ret[i] = t
	}

	return ret
}

func verticalTraversal3(root *TreeNode) [][]int {
	var res [][]int
	var position []Position
	var dfs func(*TreeNode, int, int)
	dfs = func(r *TreeNode, x int, y int) {
		if r == nil {
			return
		}
		position = append(position, Position{x, y, r.Val})
		dfs(r.Left, x+1, y-1)
		dfs(r.Right, x+1, y+1)
	}
	dfs(root, 0, 0)
	sort.Slice(position, func(i, j int) bool {
		var p1, p2 = position[i], position[j]
		if p1.y != p2.y {
			return p1.y < p2.y
		}
		if p1.x != p2.x {
			return p1.x < p2.x
		}
		return p1.val < p2.val
	})
	col := position[0].y
	arr := []int{position[0].val}
	for _, pos := range position[1:] {
		if pos.y != col {
			res = append(res, arr)
			col = pos.y
			arr = []int{pos.val}
		} else {
			arr = append(arr, pos.val)
		}
	}
	if len(arr) != 0 {
		res = append(res, arr)
	}
	return res
}

type Position struct {
	x, y, val int
}
