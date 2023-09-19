package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	// 队列走起
	var queue = make([]*TreeNode, 1, 10)
	queue[0] = root

	var res [][]int
	var t int
	var tmp []int
	var cur *TreeNode
	var layer, begin = 1, 0

	for len(queue) != 0 {
		t = len(queue)
		tmp = make([]int, t)
		if layer > 0 {
			begin = 0
		} else {
			begin = t - 1
		}
		for i := 0; i < t; i++ {
			cur = queue[i]
			tmp[begin] = cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			begin += layer
		}

		queue = queue[t:]
		layer *= -1
		res = append(res, tmp)
	}

	return res
}

func zigzagLevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var queueBuf, nextBuf [32]*TreeNode
	var queue = queueBuf[:0]
	var next = nextBuf[:0]
	queue = append(queue, root)
	var ret [][]int
	var l2r = true
	for len(queue) != 0 {
		base, add := 0, 1
		if !l2r {
			base, add = len(queue)-1, -1
		}
		var cur = make([]int, len(queue))
		for _, node := range queue {
			cur[base] = node.Val
			base += add
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		queue, next = next, queue[:0]
		ret = append(ret, cur)
		l2r = !l2r
	}
	return ret
}
