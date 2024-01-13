package main

import (
	"strconv"
	"strings"

	. "leetgo/data"
)

type Codec2 struct {
}

func Constructor2() Codec2 {
	return Codec2{}
}

// Serializes a tree to a single string.
func (c *Codec2) serialize(root *TreeNode) string {
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
func (c *Codec2) deserialize(data string) *TreeNode {
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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct{}

func Constructor5() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var sb []string
	var queue = []*TreeNode{root}

	for len(queue) != 0 {
		var cur = queue[0]
		queue = queue[1:]
		if cur == nil {
			sb = append(sb, "#")
		} else {
			sb = append(sb, strconv.Itoa(cur.Val))
			queue = append(queue, cur.Left)
			queue = append(queue, cur.Right)
		}
	}
	return strings.Join(sb, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	var sb = strings.Split(data, ",")

	// 辅助函数, 将字符串转换为一个节点
	var getNode = func(s string) *TreeNode {
		if s == "#" {
			return nil
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		return &TreeNode{Val: v}
	}

	var root = getNode(sb[0])
	if root == nil {
		return nil
	}
	var queue = []*TreeNode{root}
	var setLeft uint8
	for _, v := range sb[1:] {
		var head = queue[0]
		var node = getNode(v)
		if node != nil {
			queue = append(queue, node)
		}
		if setLeft == 0 {
			head.Left = node
		} else {
			head.Right = node
			queue = queue[1:]
		}
		setLeft ^= 1
	}
	return root
}

type Codec3 struct {
}

func Constructor4() Codec3 {
	return Codec3{}
}

// Serializes a tree to a single string.
func (c *Codec3) serialize(root *TreeNode) string {
	var sb strings.Builder

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			sb.WriteString("#,")
			return
		}
		sb.WriteString(strconv.Itoa(root.Val))
		sb.WriteByte(',')
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return sb.String()
}

// Deserializes your encoded data to tree.
func (c *Codec3) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	if len(sp) == 0 {
		return nil
	}
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "#" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{Val: val, Left: build(), Right: build()}
	}
	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
