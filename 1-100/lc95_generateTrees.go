package main

import (
	"fmt"
	. "leetgo/data"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return generate(1, n)
}

func generate(left, right int) []*TreeNode {
	if left > right {
		return []*TreeNode{nil}
	}
	res := make([]*TreeNode, 0)
	for i := left; i <= right; i++ {
		leftNodes, rightNodes := generate(left, i-1), generate(i+1, right)
		for _, lv := range leftNodes {
			for _, rv := range rightNodes {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  lv,
					Right: rv,
				})
			}
		}
	}
	return res
}

func main() {
	fmt.Println(generateTrees(3))
}
