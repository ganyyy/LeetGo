package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIteratorError struct {
	// 写一个中序遍历即可
	vals []int
	idx  int
}

func ConstructorError(root *TreeNode) BSTIteratorError {
	// 不符合题意. 因为用到的内存太多了
	var vals []int

	var iterator func(root *TreeNode)

	iterator = func(root *TreeNode) {
		if root == nil {
			return
		}
		iterator(root.Left)
		vals = append(vals, root.Val)
		iterator(root.Right)
	}

	iterator(root)

	return BSTIteratorError{
		vals: vals,
	}
}

func (bsi *BSTIteratorError) Next() int {
	bsi.idx++
	return bsi.vals[bsi.idx-1]
}

func (bsi *BSTIteratorError) HasNext() bool {
	return bsi.idx < len(bsi.vals)
}

type BSTIterator struct {
	// 写一个中序遍历即可
	nodes []*TreeNode
}

func ConstructorTrue(root *TreeNode) BSTIterator {
	var nodes []*TreeNode

	for root != nil {
		nodes = append(nodes, root)
		root = root.Left
	}
	return BSTIterator{
		nodes: nodes,
	}
}

func (bsi *BSTIterator) Next() int {
	// 栈顶一定是最小的
	var top = bsi.nodes[len(bsi.nodes)-1]
	bsi.nodes = bsi.nodes[:len(bsi.nodes)-1]

	if top.Right != nil {
		// 右边依次入栈
		var root = top.Right
		for root != nil {
			bsi.nodes = append(bsi.nodes, root)
			root = root.Left
		}
	}
	return top.Val
}

func (bsi *BSTIterator) HasNext() bool {
	return len(bsi.nodes) > 0
}

/**
 * Your BSTIteratorError object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
