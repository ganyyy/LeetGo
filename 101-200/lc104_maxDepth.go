//go:build ignore

package main

// 非递归使用层次遍历进行处理
func maxDepth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	stack := []*TreeNode{root}

	var res int
	// 理解为一层一层的来 ?
	for t := len(stack); 0 != t; t = len(stack) {
		for i := 0; i < t; i++ {
			r := stack[0]
			stack = stack[1:]
			if nil != r.Left {
				stack = append(stack, r.Left)
			}
			if nil != r.Right {
				stack = append(stack, r.Right)
			}
		}
		res++
	}
	return res
}

// 递归会更简单一点
func maxDepth2(root *TreeNode) int {
	if nil == root {
		return 0
	}
	return max(maxDepth2(root.Left), maxDepth2(root.Right)) + 1
}

func main() {

}
