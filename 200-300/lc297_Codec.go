package main

import (
	"strconv"
	"strings"

	. "leetgo/data"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var res strings.Builder
	var queue = []*TreeNode{root}
	for ln := len(queue); 0 != ln; ln = len(queue) {
		for i := 0; i < ln; i++ {
			top := queue[i]
			if nil != top {
				res.WriteString(strconv.Itoa(top.Val))
				queue = append(queue, top.Left)
				queue = append(queue, top.Right)
			} else {
				res.WriteByte('$')
			}
			res.WriteByte(',')
		}
		queue = queue[ln:]
	}
	return res.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	// 先分割字符串
	values := strings.Split(data, ",")
	// 去掉末尾的 ""
	if 0 == len(values) || "$" == values[0] {
		return nil
	}
	root := NewTreeNode(values[0])
	queue := []*TreeNode{root}
	var i int
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		if nil == cur {
			continue
		}
		cur.Left = NewTreeNode(values[i+1])
		cur.Right = NewTreeNode(values[i+2])
		i += 2
		queue = append(queue, cur.Left)
		queue = append(queue, cur.Right)
	}
	return root
}

func NewTreeNode(val string) *TreeNode {
	if val == "$" {
		return nil
	}
	v, err := strconv.Atoi(val)
	if err != nil {
		return nil
	}
	return &TreeNode{Val: v}
}
