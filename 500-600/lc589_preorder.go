package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	var ret []int

	if root == nil {
		return ret
	}

	var stack = []*Node{root}

	for len(stack) != 0 {
		var root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root == nil {
			continue
		}
		ret = append(ret, root.Val)
		for i := len(root.Children) - 1; i >= 0; i-- {
			var n = root.Children[i]
			if n == nil {
				continue
			}
			stack = append(stack, n)
		}
	}
	return ret
}
