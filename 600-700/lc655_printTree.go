package main

import "strconv"

func printTree(root *TreeNode) [][]string {

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 迭代两次?
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(height(root.Right), height(root.Left)) + 1
	}

	h := height(root)
	if h == 0 {
		return nil
	}
	var col = (1 << h) - 1
	var ret = make([][]string, h)
	for i := range ret {
		ret[i] = make([]string, col)
	}

	var fill func(*TreeNode, int, int)

	fill = func(root *TreeNode, r, c int) {
		if root == nil {
			return
		}
		ret[r][c] = strconv.Itoa(root.Val)
		if r >= h-1 {
			return
		}
		offset := 1 << (h - 1 - r - 1)
		fill(root.Left, r+1, c-offset)
		fill(root.Right, r+1, c+offset)
	}

	fill(root, 0, col/2)

	return ret
}
