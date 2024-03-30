package main

import (
	"fmt"
)

func isValidSerialization(preorder string) bool {

	// 如果遇到了 #, 说明后一个节点应该是前一个节点的子节点

	// 也是栈的一种应用吧..

	var ln = len(preorder)
	// 不满足条件的情况
	if ln == 0 || preorder[ln-1] != '#' || (ln > 1 && preorder[0] == '#') {
		return false
	}

	var stack []string

	var pre, cur int
	for cur < len(preorder) {

		switch preorder[cur] {
		case ',':
			stack = append(stack, preorder[pre:cur])
			pre = cur + 1
		case '#':
			// 如果栈为空, 那么可以直接加进去
			// 如果栈顶是数字, 也可以直接入栈
			if len(stack) == 0 || stack[len(stack)-1] != "#" {
				stack = append(stack, "#")
			} else if len(stack) >= 2 && stack[len(stack)-1] == "#" && stack[len(stack)-2] != "#" {
				// 出栈并进行替换
				// x, #, # -> #
				stack = append(stack, "#")
				for len(stack) >= 3 && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" {
					stack[len(stack)-3] = "#"
					stack = stack[:len(stack)-2]
				}
			} else {
				// 到这里的可能有:
				// 1. #, #
				return false
			}
			// 注意这里的偏移, 先去掉一个#, 同时 pre 应该从下一个字母开始
			cur++
			pre = cur + 1
		}
		cur++
	}

	// 栈中只剩余一个 # 即为合法的 二叉树 前序遍历
	return len(stack) == 1 && stack[0] == "#"

}

func isValidSerialization2(preorder string) bool {
	// 不用栈的解法

	// num 为 # 出现的次数.
	var n, num = len(preorder), 0

	for i := n - 1; i >= 0; i-- {
		switch preorder[i] {
		case ',':
			continue
		case '#':
			num++
		default:
			// 找到一个数字
			for i > 0 && preorder[i] != ',' {
				i--
			}
			// num的数量应该满足一个叶子节点, 否则就是不正确的
			if num >= 2 {
				num--
			} else {
				return false
			}
		}
	}
	// 最后只剩一个根节点
	return num == 1
}

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,#,2,#,6,#,#"))
}

func isValidSerialization3(preorder string) bool {

	const (
		Num = 0 - iota
		Nil
		End
	)

	var getNext = func() int {
		if len(preorder) == 0 {
			return End
		}
		var idx int
		for idx < len(preorder) && preorder[idx] != ',' {
			idx++
		}
		var ret = Num
		if preorder[0] == '#' {
			ret = Nil
		}
		// fmt.Println(ret, preorder)
		// 跳过','
		if idx == len(preorder) {
			idx--
		}
		preorder = preorder[idx+1:]
		return ret
	}

	var build func() bool
	build = func() bool {
		next := getNext()
		if next == Nil {
			return true
		}
		if next == End {
			// 不允许提前结束?
			return false
		}

		// 左右分支
		leftValid := build()
		if !leftValid {
			return false
		}
		return build()
		// fmt.Println(next, leftValid, rightValid)
		// if !leftValid || !rightValid {
		//     return false
		// }
		// return true
	}

	// 全部消耗完毕
	return build() && len(preorder) == 0
}

func isValidSerializationStack(preorder string) bool {

	const (
		Num = 0 - iota
		Nil
		End
	)

	var getNext = func() int {
		if len(preorder) == 0 {
			return End
		}
		var idx int
		for idx < len(preorder) && preorder[idx] != ',' {
			idx++
		}
		var ret = Num
		if preorder[0] == '#' {
			ret = Nil
		}
		// fmt.Println(ret, preorder)
		// 跳过','
		if idx == len(preorder) {
			idx--
		}
		preorder = preorder[idx+1:]
		return ret
	}

	var stack []int

	for next := getNext(); next != End; next = getNext() {
		stack = append(stack, next)
		for len(stack) >= 3 && stack[len(stack)-1] == Nil && stack[len(stack)-2] == Nil && stack[len(stack)-3] != Nil {
			stack = stack[:len(stack)-3]
			stack = append(stack, Nil)
		}
	}
	return len(stack) == 1 && stack[0] == Nil
}

func isValidSerializationDegree(preorder string) bool {

	const (
		Num = 0 - iota
		Nil
		End
	)

	var getNext = func() int {
		if len(preorder) == 0 {
			return End
		}
		var idx int
		for idx < len(preorder) && preorder[idx] != ',' {
			idx++
		}
		var ret = Num
		if preorder[0] == '#' {
			ret = Nil
		}
		// fmt.Println(ret, preorder)
		// 跳过','
		if idx == len(preorder) {
			idx--
		}
		preorder = preorder[idx+1:]
		return ret
	}

	var diff int // 出度-入读
	// 根节点出度是2, 入度是0
	// 非根节点的出度是2, 入度是1
	// 叶子节点的出度是0, 入度是1
	// 整棵树迭代完成后, 其差值应该是0!
	diff++
	for next := getNext(); next != End; next = getNext() {
		diff--
		if diff < 0 {
			return false
		}
		if next != Nil {
			diff += 2
		}
	}
	return diff == 0
}
