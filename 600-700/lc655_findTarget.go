package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	var stack []*TreeNode
	var vals = make(map[int]bool)

	for root != nil || len(stack) != 0 {
		if root == nil {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		} else {
			if vals[k-root.Val] {
				return true
			}
			vals[root.Val] = true
			stack = append(stack, root)
			root = root.Left
		}
	}
	// fmt.Println(vals)
	return false
}
