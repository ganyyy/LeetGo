package main

/**
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

这里采用的方法是以空间换时间.
*/

func longestConsecutive(nums []int) int {
	// 每个端点值对应的连续区间长度
	m := make(map[int]int, len(nums))
	var res int
	for _, v := range nums {
		// 如果已存在, 直接略过
		if _, ok := m[v]; ok {
			continue
		}
		// 取左右的长度
		l, r := m[v-1], m[v+1]
		// 更新最大长度
		var now = l + r + 1
		if now > res {
			res = now
		}
		// 更新该连续区间左右端点的值
		// 如果后续这两个端点可以继续延伸, 那么就可以直接通过这两个端点来获取最大长度
		m[v] = now
		m[v-l] = now
		m[v+r] = now

	}

	return res
}

// 定义一个并查集

type UF struct {
	Parent []int // 对应位置的根节点
	Size   []int // 对应位置的子节点数目, 当且仅当Parent[v] == v时有效
	Count  int
}

func NewUF(n int) UF {
	var uf UF
	uf.Parent = make([]int, n)
	uf.Size = make([]int, n)
	for i := range uf.Parent {
		uf.Parent[i] = i // 指向自己
		uf.Size[i] = 1   // 权重为1
	}
	uf.Count = n
	return uf
}

func (u *UF) Find(v int) int {
	// 压平,
	// 1,2,3,4,5
	// [1,1,2,3,4] -> [1,1,1,1,1]
	if u.Parent[v] != v {
		u.Parent[v] = u.Find(u.Parent[v])
	}
	return u.Parent[v]
}

func (u *UF) Union(a, b int) {
	var pa, pb = u.Find(a), u.Find(b)
	if pa == pb {
		return
	}
	// pb始终是Size较大的一方, 防止整体退化成单链表
	// 将pa挂到pb上
	if u.Size[pa] > u.Size[pb] {
		pa, pb = pb, pa
	}
	u.Parent[pa] = pb
	u.Size[pb] += u.Size[pa]
	u.Count--
}

func (u *UF) Connected(a, b int) bool {
	return u.Find(a) == u.Find(b)
}

func longestConsecutive2(nums []int) int {
	var uf = NewUF(len(nums))
	var m = make(map[int]int, len(nums))

	for i, v := range nums {
		if _, ok := m[v]; ok {
			continue
		}
		if j, ok := m[v-1]; ok {
			uf.Union(i, j)
		}
		if j, ok := m[v+1]; ok {
			uf.Union(i, j)
		}
		m[v] = i
	}

	// 找到最大的联通分量
	var mx int
	for i, v := range uf.Parent {
		if i == v {
			if mx < uf.Size[i] {
				mx = uf.Size[i]
			}
		}
	}
	return mx
}
