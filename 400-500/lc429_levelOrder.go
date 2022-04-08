package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	var queue []*Node

	queue = append(queue, root)
	for len(queue) != 0 {
		var ln = len(queue)
		var arr = make([]int, 0, ln)
		for i := 0; i < ln; i++ {
			var node = queue[i]
			arr = append(arr, node.Val)

			for _, child := range node.Children {
				if child == nil {
					continue
				}
				queue = append(queue, child)
			}
		}
		ret = append(ret, arr)
		queue = queue[ln:]
	}

	return ret
}
