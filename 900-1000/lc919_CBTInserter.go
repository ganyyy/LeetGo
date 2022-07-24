//go:build ignore

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	Root  *TreeNode
	Queue []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	// 层次遍历, 获取所有的子节点信息
	if root == nil {
		return CBTInserter{}
	}
	var queue []*TreeNode
	var allNode = []*TreeNode{nil} // 填充一个dummy节点
	var add = func(node *TreeNode) {
		if node == nil {
			return
		}
		queue = append(queue, node)
	}
	queue = append(queue, root)
	for len(queue) != 0 {
		var ln = len(queue)
		for i := 0; i < ln; i++ {
			var cur = queue[i]
			allNode = append(allNode, cur)
			add(cur.Left)
			add(cur.Right)
		}
		queue = queue[ln:]
	}
	return CBTInserter{
		Root:  root,
		Queue: allNode,
	}
}

func (this *CBTInserter) Insert(val int) int {
	var node = &TreeNode{
		Val: val,
	}
	if this.Root == nil {
		this.Root = node
		this.Queue = append(this.Queue, node)
		return 0
	}
	// TODO
	// 计算出当前的位置何其对应的父节点位置
	var idx = len(this.Queue)
	this.Queue = append(this.Queue, node)
	var parent = this.Queue[idx/2]
	if idx&1 == 0 {
		// 偶数, 在左子节点上
		parent.Left = node
	} else {
		parent.Right = node
	}
	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.Root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
