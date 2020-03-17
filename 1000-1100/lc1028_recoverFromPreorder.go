package main

import (
	. "leetgo/data"
	"strconv"
)

func recoverFromPreorder(S string) *TreeNode {
	return build([]byte(S))
}

func build(s []byte) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	t := new(TreeNode)
	start := 0
	for ; start < len(s) && s[start] >= '0' && s[start] <= '9'; start++ {
	}
	t.Val, _ = strconv.Atoi(string(s[:start]))
	s = s[start:]
	// 如果后边还有跟着的, 就一定是左孩子, 由此可以计算出是第几层, 并由计算得来的深度确定右孩子存不存在
	for start = 0; start < len(s) && s[start] == '-'; start++ {
	}
	if start == 0 {
		return t
	}
	s = s[start:]
	rIndex := -1
	for i, cur := 0, 0; i < len(s); i++ {
		if s[i] == '-' {
			cur++
			// 当前为-, 下一个不是- 并且和当前深度相等
			if s[i+1] != '-' && cur == start {
				// 右节点找到了
				rIndex = i + 1
				break
			}
		} else {
			cur = 0
		}
	}
	if rIndex != -1 {
		t.Left = build(s[:rIndex-start])
		t.Right = build(s[rIndex:])
	} else {
		t.Left = build(s)
	}
	return t
}

func main() {
	recoverFromPreorder("1-2--3--4-5--6")
}
