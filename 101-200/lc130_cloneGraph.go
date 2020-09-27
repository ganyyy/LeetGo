//+build ignore

package main

type Node struct {
	Val       int
	Neighbors []*Node
}
type M map[int]*Node

func cloneGraph(node *Node) *Node {
	if nil == node {
		return nil
	}
	m := make(map[int]*Node)
	return dfs(node, m)
}

func dfs(n *Node, m M) *Node {
	if d, ok := m[n.Val]; ok {
		return d
	}
	v := &Node{Val: n.Val}
	m[v.Val] = v
	for _, t := range n.Neighbors {
		if d, ok := m[t.Val]; ok {
			v.Neighbors = append(v.Neighbors, d)
		} else {
			v.Neighbors = append(v.Neighbors, dfs(t, m))
		}
	}
	return v
}
