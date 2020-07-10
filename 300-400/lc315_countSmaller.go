package main

import "fmt"

type Node struct {
	// 当前节点的值, 当前节点的左子树的节点个数
	val, cnt    int
	left, right *Node
}

func newNode(val int) *Node {
	return &Node{
		val: val,
	}
}

func (n *Node) insert(val int, res *int) {
	if n.val >= val {
		// 如果值比当前值要小, 当前节点的左孩子节点个数+1
		n.cnt++
		if nil == n.left {
			n.left = newNode(val)
		} else {
			n.left.insert(val, res)
		}
	} else {
		// 找到了比新值小的节点, 就加上比当前节点的cnt以及 当前节点本身
		*res += n.cnt + 1
		if nil == n.right {
			n.right = newNode(val)
		} else {
			n.right.insert(val, res)
		}
	}
}

func countSmaller(nums []int) []int {
	ln := len(nums)
	if 0 == ln {
		return nums
	}
	res := make([]int, ln)
	root := newNode(nums[ln-1])

	for i := ln - 2; i >= 0; i-- {
		root.insert(nums[i], &res[i])
	}
	return res
}

func main() {
	fmt.Println(countSmaller([]int{5, 2, 6, 1}))
}
