//go:build ignore

package main

const kLog = 16

type TreeAncestor struct {
	ancestors [][kLog]int
}

func Constructor(n int, parent []int) TreeAncestor {
	var ta TreeAncestor
	// 节点i的第2^j个祖先
	// 我们需要存储的父节点分别是: 0,1,2,4,8,16 ..., 所以5w个节点的最大层级是16
	// 如果要查找的是二的整数次幂的父节点, 可以直接找到数组的最高位1定位
	// 那么: 如何应对非二的整数次幂的父节点呢?
	// 这里面就存在了状态转移问题.
	// 比如第3个父节点, 其二进制表示为 0b11,
	// 第一位不为0, 则跳转到父节点, 此时就变成了查找父节点的第二个父节点,
	// 比如第6个父节点, 其二进制表示为 0b110,
	// 第一位是0, 不需要操作, 第二位不为零, 则先跳转到第二个父节点, 与之对应的就是第二个父节点的第四个父节点
	ta.ancestors = make([][kLog]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < kLog; j++ {
			ta.ancestors[i][j] = -1
		}
		ta.ancestors[i][0] = parent[i]
	}
	for j := 1; j < kLog; j++ {
		for i := 0; i < n; i++ {
			if ta.ancestors[i][j-1] != -1 {
				// j = 0, 当前节点的第一个祖先, 也就是父节点
				// j = 1, 当前节点的第二个祖先, 也就是父节点的第一个祖先
				// j = 2, 当前节点的第四个祖先, 也就是第二个父节点的第二个祖先
				// j = 3, 当前节点的第八个祖先, 也就是第四个父节点的第四个祖先

				// 因为外围的循环是j, 所以这里的j-1就是上一层的祖先, 所以当迭代到j的时候, 所有小于j的各个节点的祖先都已经计算完毕
				ta.ancestors[i][j] = ta.ancestors[ta.ancestors[i][j-1]][j-1]
			}
		}
	}
	return ta
}

func (ta *TreeAncestor) GetKthAncestor(node int, k int) int {
	for j := 0; j < kLog; j++ {
		if (k>>j)&1 != 0 {
			node = ta.ancestors[node][j]
			if node == -1 {
				return -1
			}
		}
	}
	return node
}
