package main

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	var dfs func(preorder, postorder []int) *TreeNode
	//   1
	// 2   3
	// 4 5 6 7
	// pre  [1,2,4,5,3,6,7]
	// post [4,5,2,6,7,3,1]
	dfs = func(preorder, postorder []int) *TreeNode {
		if len(preorder) == 0 {
			return nil
		}

		root := preorder[0]
		// 去根节点
		preorder = preorder[1:]
		postorder = postorder[:len(postorder)-1]

		if len(preorder) == 0 {
			return &TreeNode{Val: root}
		}

		// 左子树的长度
		var leftCount int
		if len(preorder) > 0 {
			leftRoot := preorder[0]
			for ; leftCount < len(postorder); leftCount++ {
				if postorder[leftCount] == leftRoot {
					break
				}
			}
		}

		// 分割
		return &TreeNode{
			Val: root,
			// [2,4,5], [4,5,2]
			Left: dfs(preorder[:leftCount+1], postorder[:leftCount+1]),
			// [3,6,7], [6,7,3]
			Right: dfs(preorder[leftCount+1:], postorder[leftCount+1:]),
		}
	}
	return dfs(preorder, postorder)
}
