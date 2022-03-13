package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	// 层次遍历啊
	var queue = []*TreeNode{root}
	var ret []int
	for len(queue) != 0 {
		var top = queue[0]
		queue = queue[1:]
		ret = append(ret, top.Val)
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
	return ret
}

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var queue = []*TreeNode{root}
	var ret [][]int
	var add = func(node *TreeNode) {
		if node == nil {
			return
		}
		queue = append(queue, node)
	}
	for len(queue) != 0 {
		var ln = len(queue)
		var arr = make([]int, 0, ln)
		for i := 0; i < ln; i++ {
			var n = queue[i]
			arr = append(arr, n.Val)
			add(n.Left)
			add(n.Right)
		}
		ret = append(ret, arr)
		queue = queue[ln:]
	}
	return ret
}

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var queue = []*TreeNode{root}
	var ret [][]int
	var add = func(node *TreeNode) {
		if node == nil {
			return
		}
		queue = append(queue, node)
	}

	var reverse = func(arr []int) {
		for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}

	var cnt int

	for len(queue) != 0 {
		var ln = len(queue)
		var arr = make([]int, 0, ln)
		for i := 0; i < ln; i++ {
			var n = queue[i]
			arr = append(arr, n.Val)
			add(n.Left)
			add(n.Right)
		}
		cnt++
		if cnt&1 == 0 {
			reverse(arr)
		}
		ret = append(ret, arr)
		queue = queue[ln:]
	}

	// for i := 1; i < len(ret); i += 2 {
	//     reverse(ret[i])
	// }

	return ret
}
