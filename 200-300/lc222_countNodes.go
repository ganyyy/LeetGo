package main

import . "leetgo/data"

func countNodes(root *TreeNode) int {
	if nil == root {
		return 0
	}
	// 只需要确定子树的高度(左孩子节点的层级)就可以了
	ld := getDepth(root.Left)
	rd := getDepth(root.Right)
	// 左右子树的层级相同, 说明右边最后一层不满, 递归的统计右边的节点个数
	if ld == rd {
		// 计算节点个数, ld层一共  2^ld - 1 + 1(根节点数量) = 1 << ld
		return countNodes(root.Right) + (1 << ld)
	} else {
		// 左右层级不同, 说明左边最后一层不满, 右边倒数第二层是满的, 递归的统计左子树即可
		return countNodes(root.Left) + (1 << rd)
	}
}

func getDepth(root *TreeNode) int {
	var res int
	for root != nil {
		res++
		root = root.Left
	}
	return res
}

func main() {

}
