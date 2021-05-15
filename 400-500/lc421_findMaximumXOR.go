package main

type node struct {
	next [2]*node
}

func newNode() *node {
	return &node{}
}

func insert(num int, root *node) {
	// 从高位开始看, 因为高位的优先级高于低位
	for i := 30; i >= 0; i-- {
		var t = (num >> i) & 1
		if root.next[t] == nil {
			root.next[t] = newNode()
		}
		root = root.next[t]
	}
}

func findMaximumXOR(nums []int) int {
	var root = newNode()
	for _, v := range nums {
		insert(v, root)
	}

	var res int
	for _, v := range nums {
		var p = root
		var tmp int
		for i := 30; i >= 0; i-- {
			var t = (v >> i) & 1
			// 为了能让当前位异或的结果为1
			// 同样的, 这里的每一条路线都可以走到末尾
			// 按照贪心原则, 每一位都选取最大收益, 那么最终结果就是当前数可以匹配的最大异或值
			if p.next[t^1] != nil {
				tmp += 1 << i
				p = p.next[t^1]
			} else {
				p = p.next[t]
			}
		}
		if tmp > res {
			res = tmp
		}
	}

	return res
}
