package main

const (
	BitCount = 15 // 根据数据范围, 确定的位数
)

// Node 二维字典树
type Node struct {
	Next [2]*Node
	Sum  int
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) Add(num int) {
	cur := n
	for i := BitCount; i >= 0; i-- {
		bit := (num >> i) & 1

		next := cur.Next[bit]
		if next == nil {
			next = NewNode()
			cur.Next[bit] = next
		}
		cur = next
		cur.Sum++
	}
}

func (n *Node) Get(num, x int) (sum int) {
	// 从字典树中, 找到存在多少个数字, 满足 ai⊕num <= x
	cur := n
	for i := BitCount; i >= 0; i-- {
		if cur == nil {
			break
		}
		bit := (num >> i) & 1
		xBit := (x >> i) & 1

		if xBit != 0 {
			// x[i] == 1

			next := cur.Next[bit]
			if next != nil {
				sum += next.Sum
			}

			// 异或后相等
			// num[i] == 1, ax[i] = 0
			// num[i] == 0, ax[i] = 1
			cur = cur.Next[bit^1]
		} else {
			// x[i] == 0, 则不会存在任何一个 ai⊕num < x
			cur = cur.Next[bit]
		}
	}
	return
}

func countPairs(nums []int, low int, high int) int {
	var root Node
	var ret int
	for _, v := range nums[:] {
		ret += root.Get(v, high+1) - root.Get(v, low)
		root.Add(v)
	}

	return ret
}
