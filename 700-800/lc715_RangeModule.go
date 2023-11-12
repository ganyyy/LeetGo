package main

import "github.com/emirpasic/gods/trees/redblacktree"

type RangeModule struct {
	*redblacktree.Tree
}

func Constructor715() RangeModule {
	return RangeModule{redblacktree.NewWithIntComparator()}
}

func (t RangeModule) AddRange(left, right int) {
	// 每次AddRange, 都需要保证树中存在的区间无法进一步合并
	// key表示左区间, value表示右区间

	// 通过Floor找出 node.Key <= left的最大值
	if node, ok := t.Floor(left); ok {
		r := node.Value.(int)
		if r >= right {
			// node.Value >= right, 说明 [left, right) 已包含在树中
			// 无需任何操作
			return
		}
		if r >= left {
			// 如果两个区间存在重叠 left1 <= left <= right1 <= right
			// 就需要将两个区间进行合并
			// 这里先删除一下旧的key, 最后再插进去
			left = node.Key.(int)
			t.Remove(left)
		}
	}
	// 通过Ceil找出 node.Key >= left的最小值
	for node, ok := t.Ceiling(left); ok && node.Key.(int) <= right; node, ok = t.Ceiling(left) {
		// 还是合并, 不停的合并 [left, right) 间的区间
		right = max(right, node.Value.(int))
		// 合并完的就删去
		t.Remove(node.Key)
	}
	// 将最后计算完成的区间再次添加进去
	t.Put(left, right)
}

func (t RangeModule) QueryRange(left, right int) bool {
	// 符合条件的区间, 就是 node.Key <= left < right <= node.Value
	node, ok := t.Floor(left)
	return ok && node.Value.(int) >= right
}

func (t RangeModule) RemoveRange(left, right int) {
	if node, ok := t.Floor(left); ok {
		// 找前置节点, 这次的分割可能会将之前的一个区间分割成两个
		l, r := node.Key.(int), node.Value.(int)
		if r >= right {
			// node.Key <= left <= right <= node.Value =>
			// 	[node.key, left) + [right, node.Value)
			if l == left {
				t.Remove(l)
			} else {
				node.Value = left
			}
			if right != r {
				t.Put(right, r)
			}
			return
		}
		if r > left {
			// node.Key <= left < node.Value < right =>
			// 	[node.key, left) + ...
			node.Value = left
		}
	}

	// 通过Ceil找出 node.Key >= left的最小值
	for node, ok := t.Ceiling(left); ok && node.Key.(int) < right; node, ok = t.Ceiling(left) {
		// 移除所有在 [left, right) 中的所有区间
		//	... + [node.Key, node.Value) + ...
		r := node.Value.(int)
		t.Remove(node.Key)
		if r > right {
			// 找到最后一个区间, 添加进去
			// [right, r)
			t.Put(right, r)
			break
		}
	}
}
