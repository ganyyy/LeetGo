package main

func flatten(root *TreeNode) {
	if nil == root {
		return
	}
	// 层次遍历解开即可
	var stack []*TreeNode
	var res []*TreeNode
	cur := root
	for len(stack) != 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = cur.Right
		} else {
			stack = append(stack, cur)
			res = append(res, cur)
			cur = cur.Left
		}
	}
	if len(res) == 0 {
		return
	}
	pre := res[0]
	for i := 1; i < len(res); i++ {
		pre.Left = nil
		pre.Right = res[i]
		pre = pre.Right
	}
}

func flatten2(root *TreeNode) {

	/*

		比如这棵树. 1是根节点, 2是左子节点, 5是右子节点.
		查找5的前驱节点, 也就是2的最右子节点, 也就是4
		随后将 4的右子节点指向5, 将1的右子节点指向2, 将1的左子节点置空

		       1
		    2     5
		   3 4      6

		依次迭代每个右子节点, 直到为空
		        1
		          2
		        3   4
		              5
		                6

		        1
		          2
		            3
		              4
		                5
		                  6
	*/

	curr := root
	for curr != nil {
		// 如果当前节点的左子节点不为空, 就找到右子节点的前驱节点
		// 这个右子节点的前驱, 就是左子节点一直往右找到空的位置
		// 然后将右子节点接到前驱上, 将左子节点转移到右子节点, 实现一次旋转(!)
		if curr.Left != nil {
			next := curr.Left
			predecessor := next
			for predecessor.Right != nil {
				predecessor = predecessor.Right
			}
			// 前驱节点找到后, 将 当前节点的右节点赋给 前驱节点的右节点
			predecessor.Right = curr.Right
			// 当前节点的左节点清空, 右键点指向左节点
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}

func main() {
	t := &TreeNode{Val: 1}
	t.Left = &TreeNode{Val: 2}
	t.Right = &TreeNode{Val: 5}
	t.Left.Left = &TreeNode{Val: 3}
	t.Left.Right = &TreeNode{Val: 4}
	t.Right.Left = &TreeNode{Val: 6}

	flatten(t)
}
