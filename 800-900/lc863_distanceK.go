package main

import . "leetgo/data"

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var parentMap = make(map[int]*TreeNode, 0)
	var visited = make(map[int]bool, 0)

	var dfs func(root *TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			parentMap[root.Left.Val] = root
			dfs(root.Left)
		}
		if root.Right != nil {
			parentMap[root.Right.Val] = root
			dfs(root.Right)
		}
	}

	dfs(root)

	var ret []int

	var dfs2 func(root *TreeNode, cnt int)

	dfs2 = func(root *TreeNode, cnt int) {
		// fmt.Println(root, cnt)
		if root == nil || cnt < 0 {
			return
		}
		if visited[root.Val] {
			return
		}
		visited[root.Val] = true
		if cnt == 0 {
			ret = append(ret, root.Val)
		} else if cnt > 0 {
			cnt--
			dfs2(root.Left, cnt)
			dfs2(root.Right, cnt)
			dfs2(parentMap[root.Val], cnt)
		}
	}

	dfs2(target, k)

	return ret
}

const (
	// 指定了下一步遍历的方向

	LL  int = 1
	RR  int = 2
	ALL int = 0
)

// 瞻仰一下大佬的操作
func distanceK2(root *TreeNode, target *TreeNode, K int) []int {

	var res []int
	var find func(root *TreeNode, k int, lr int)
	find = func(root *TreeNode, k int, lr int) {
		if root == nil || k < 0 || k != K && root == target {
			return
		}
		if k == 0 {
			res = append(res, root.Val)
			return
		}
		if lr == LL || lr == ALL {
			find(root.Left, k-1, ALL)
		}
		if lr == RR || lr == ALL {
			find(root.Right, k-1, ALL)
		}
	}

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return -1
		}

		// 根节点, 左右一起找
		if root == target {
			find(root, K, ALL)
			return K - 1
		}

		// 查询左子树, 从root的右边开始找
		l := dfs(root.Left)
		if l >= 0 {
			find(root, l, RR)
			return l - 1
		}

		// 查询右子树, 从root的左边开始找
		r := dfs(root.Right)
		if r >= 0 {
			find(root, r, LL)
			return r - 1
		}

		return -1
	}

	dfs(root)

	return res
}
