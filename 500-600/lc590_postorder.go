//go:build ignore
// +build ignore

package main

type Node struct {
	Val      int
	Children []*Node
}

// func postorder(root *Node) []int {
//     if root == nil {
//         return nil
//     }
//     var ret []int
//     for _, p := range root.Children {
//         ret = append(ret, postorder(p)...)
//     }
//     ret = append(ret, root.Val)
//     return ret
// }

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	var nodeStack []*Node
	var outputStack []int
	nodeStack = append(nodeStack, root)

	for len(nodeStack) != 0 {
		var node = nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		outputStack = append(outputStack, node.Val)

		for _, p := range node.Children {
			if p == nil {
				continue
			}
			nodeStack = append(nodeStack, p)
		}
	}

	for l, r := 0, len(outputStack)-1; l < r; l, r = l+1, r-1 {
		outputStack[l], outputStack[r] = outputStack[r], outputStack[l]
	}

	return outputStack
}
