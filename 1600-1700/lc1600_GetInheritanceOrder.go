package main

//
//type node struct {
//	// 当前节点的名字指针
//	name *string
//	// 父节点名字指针
//	parent *string
//	// 兄弟节点
//	prev, next *node
//	// 子节点
//	childes *node
//	// 最后一个子节点, 方便执行插入
//	lastChild *node
//}
//
//func newNode(name string) *node {
//	var n = new(string)
//	*n = name
//	return &node{
//		name: n,
//	}
//}
//
//type ThroneInheritance struct {
//	// 分层
//	root *node
//	// 快速定位
//	m map[string]*node
//	// 整体是一个树状结构, 新增就是在指定节点下添加新的子节点
//
//	cnt int
//}
//
//func Constructor(kingName string) ThroneInheritance {
//	var t = ThroneInheritance{}
//
//	t.m = map[string]*node{}
//
//	t.root = newNode(kingName)
//
//	t.m[kingName] = t.root
//
//	return t
//}
//
//func (ti *ThroneInheritance) Birth(parentName string, childName string) {
//	var n = ti.m[parentName]
//	if n == nil {
//		return
//	}
//	var child = newNode(childName)
//	child.parent = n.name
//	if n.lastChild != nil {
//		child.prev = n.lastChild
//		child.next = n.lastChild.next
//		if child.next != nil {
//			child.next.prev = child
//		}
//		n.lastChild.next = child
//	} else {
//		// 为空有两种可能
//		if n.childes == nil {
//			// 1. 没有任何孩子节点
//			n.childes = child
//		} else {
//			// 2. 存在原来的兄弟节点
//			n.childes.prev = child
//			child.next = n.childes
//		}
//	}
//	n.lastChild = child
//
//	ti.m[childName] = child
//
//	ti.cnt++
//}
//
//func (ti *ThroneInheritance) Death(name string) {
//	var n = ti.m[name]
//	if n == nil {
//		return
//	}
//	delete(ti.m, name)
//
//	// 先看有没有后代
//	if n.childes != nil {
//		// 如果存在后代的话, 后代直接继承父亲位置
//		var child = n.childes
//		n.childes = n.childes.next
//		if n.childes != nil {
//			n.childes.prev = nil
//		}
//
//		// 将直系后代的字节点末尾节点串联上 旁系后代
//		if child.lastChild != nil {
//			child.lastChild.next = child.next
//		} else {
//			child.childes = child.next
//		}
//		if child.next != nil {
//			child.next.prev = child.lastChild
//		}
//
//		n.childes = child.childes
//		n.lastChild = child.lastChild
//		n.name = child.name
//		// 构建新的映射
//		ti.m[*n.name] = n
//	} else {
//		// 删除时也有多种可能
//		if n.prev == nil {
//			// 删除的是头节点
//			if n.parent != nil {
//				var parent = ti.m[*n.parent]
//				parent.childes = n.next
//				// 如果删除的是直系后代的最后一个, 直接清空即可
//				if n == parent.lastChild {
//					parent.lastChild = nil
//				}
//			}
//		} else {
//			// 删除的是中间节点
//			n.prev.next = n.next
//			if n.next != nil {
//				n.next.prev = n.prev
//			}
//			if n.parent != nil {
//				var parent = ti.m[*n.parent]
//				// 如果是删除的直系后代的最后一个, 需要将其意向前方
//				if parent.lastChild == n {
//					parent.lastChild = n.prev
//				}
//			}
//		}
//	}
//
//	ti.cnt--
//}
//
//func (ti *ThroneInheritance) GetInheritanceOrder() []string {
//	var ret = make([]string, 0, ti.cnt)
//	// 先序遍历
//
//	var front func(n *node)
//
//	front = func(n *node) {
//		if n == nil {
//			return
//		}
//		ret = append(ret, *n.name)
//		for t := n.childes; t != nil; t = t.next {
//			front(t)
//		}
//	}
//
//	front(ti.root)
//
//	return ret
//}

type ThroneInheritance struct {
	kingName string

	childes map[string][]string

	death map[string]bool

	cnt int
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		kingName: kingName,
		childes:  map[string][]string{},
		death:    map[string]bool{},
	}
}

func (ti *ThroneInheritance) Birth(parentName string, childName string) {
	ti.childes[parentName] = append(ti.childes[parentName], childName)
	ti.cnt++
}

func (ti *ThroneInheritance) Death(name string) {
	ti.death[name] = true
	ti.cnt--
}

func (ti *ThroneInheritance) GetInheritanceOrder() []string {
	var ret = make([]string, 0, ti.cnt)
	var dfs func(name string)

	dfs = func(name string) {
		if !ti.death[name] {
			ret = append(ret, name)
		}
		for _, n := range ti.childes[name] {
			dfs(n)
		}
	}

	dfs(ti.kingName)

	return ret
}

func main() {

}
