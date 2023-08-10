package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	var ret int
	for _, child := range root.Children {
		ret = max(ret, maxDepth(child))
	}
	return ret + 1
}
