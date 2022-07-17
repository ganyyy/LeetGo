//go:build ignore

package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func intersect(quadTree1, quadTree2 *Node) *Node {
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			// 如果一方是叶子节点, 那么以叶子节点的值为准
			return &Node{Val: true, IsLeaf: true}
		}
		return quadTree2
	}
	if quadTree2.IsLeaf {
		// 这一步相当于取巧重复使用上边的判断逻辑
		return intersect(quadTree2, quadTree1)
	}
	// 合并子节点
	o1 := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	o2 := intersect(quadTree1.TopRight, quadTree2.TopRight)
	o3 := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	o4 := intersect(quadTree1.BottomRight, quadTree2.BottomRight)
	if o1.IsLeaf && o2.IsLeaf && o3.IsLeaf && o4.IsLeaf && o1.Val == o2.Val && o1.Val == o3.Val && o1.Val == o4.Val {
		// 四个子节点全部都是叶子节点, 且值相同(0/1)
		return &Node{Val: o1.Val, IsLeaf: true}
	}
	// 不满足上述条件, 当前节点是一个非叶子节点
	return &Node{false, false, o1, o2, o3, o4}
}
