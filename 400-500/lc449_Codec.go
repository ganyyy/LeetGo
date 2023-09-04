//go:build ignore

package main

import (
	. "leetgo/data"
	"math"
	"strconv"
	"strings"
)

type Codec struct{}

func Constructor() (_ Codec) { return }

func (Codec) serialize(root *TreeNode) string {
	// mark
	var arr []string
	var postOrder func(*TreeNode)
	postOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		arr = append(arr, strconv.Itoa(node.Val))
	}
	postOrder(root)
	return strings.Join(arr, " ")
}

func (Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, " ")
	var construct func(int, int) *TreeNode
	construct = func(lower, upper int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(arr[len(arr)-1])
		if val < lower || val > upper {
			return nil
		}
		arr = arr[:len(arr)-1]
		// 怎么就行了呢?
		// 因为整体的 arr是一直被消耗的
		// 这也是为啥先迭代右边, 再迭代左边的原因
		return &TreeNode{Val: val, Right: construct(val, upper), Left: construct(lower, val)}
	}
	return construct(math.MinInt32, math.MaxInt32)
}
