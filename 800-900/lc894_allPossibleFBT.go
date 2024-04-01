package main

func allPossibleFBT(n int) []*TreeNode {
	var fullBinaryTrees []*TreeNode
	if n%2 == 0 {
		return fullBinaryTrees
	}
	if n == 1 {
		fullBinaryTrees = append(fullBinaryTrees, &TreeNode{Val: 0})
		return fullBinaryTrees
	}
	// 左右互补, 其总和为n-1, 且左右子树的节点数必须是奇数
	for i := 1; i < n; i += 2 {
		leftSubtrees := allPossibleFBT(i)
		rightSubtrees := allPossibleFBT(n - 1 - i)
		for _, leftSubtree := range leftSubtrees {
			for _, rightSubtree := range rightSubtrees {
				root := &TreeNode{Val: 0, Left: leftSubtree, Right: rightSubtree}
				fullBinaryTrees = append(fullBinaryTrees, root)
			}
		}
	}
	return fullBinaryTrees
}
