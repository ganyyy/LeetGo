package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIteratorError struct {
	// mark
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
	// 相当于将中序遍历的结果保存下来, 整体拆成了两个步骤
	// 中序遍历的顺序是左中右.
	// 一直找到最左边的节点, 然后依次出栈, 出栈的时候, 如果有右节点, 就将右节点的左节点依次入栈
	nodes []*TreeNode
}

func ConstructorTrue(root *TreeNode) BSTIterator {
	return BSTIterator{
		nodes: flattenLeft(root),
	}
}

func flattenLeft(root *TreeNode) []*TreeNode {
	var nodes []*TreeNode
	for root != nil {
		nodes = append(nodes, root)
		root = root.Left
	}
	return nodes
}

func (bsi *BSTIterator) Next() int {
	// 栈顶一定是最小的
	var top = bsi.nodes[len(bsi.nodes)-1]
	bsi.nodes = bsi.nodes[:len(bsi.nodes)-1]

	if top.Right != nil {
		// 右边依次入栈
		bsi.nodes = append(bsi.nodes, flattenLeft(top.Right)...)
	}
	return top.Val
}

func (bsi *BSTIterator) HasNext() bool {
	return len(bsi.nodes) > 0
}

/**
 * Your BSTIteratorError object will be instantiated and called as such:
 * obj := Constructor146_2(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
