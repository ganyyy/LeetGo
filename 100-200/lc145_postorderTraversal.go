//go:build ignore
// +build ignore

package main

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// 一个节点存放栈, 一个输出栈
	var nodeStack, outputStack []*TreeNode
	res := make([]int, 0)

	nodeStack = append(nodeStack, root)
	for len(nodeStack) != 0 {
		t := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		outputStack = append(outputStack, t)

		// 注意, 先左边入栈, 然后右边入栈, 这样右边会先出栈,
		// 对应在 输出栈中就是  左右中  的顺序
		if t.Left != nil {
			nodeStack = append(nodeStack, t.Left)
		}
		if t.Right != nil {
			nodeStack = append(nodeStack, t.Right)
		}
	}

	for i := len(outputStack) - 1; i >= 0; i-- {
		res = append(res, outputStack[i].Val)
	}

	return res
}

func main() {

}
