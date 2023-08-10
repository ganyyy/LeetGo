//go:build ignore

package main

type pair struct{ start, end int }
type MyCalendarTwo struct{ booked, overlaps []pair }

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (c *MyCalendarTwo) Book(start, end int) bool {
	// 已经双重预定的区间
	for _, p := range c.overlaps {
		if p.start < end && start < p.end {
			return false
		}
	}
	// 首次预定的区间
	for _, p := range c.booked {
		if p.start < end && start < p.end {
			c.overlaps = append(c.overlaps, pair{max(p.start, start), min(p.end, end)})
		}
	}
	c.booked = append(c.booked, pair{start, end})
	return true
}

type pair struct{ first, second int }
type MyCalendarTwo map[int]pair

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (tree MyCalendarTwo) update(start, end, val, l, r, idx int) {
	// start: 更新的左区间
	//  end : 更新的右区间
	//  val :   更新的值
	//   l  : 当前所在的左区间
	//   r  : 当前所在的右区间
	//  idx : 当前要更新的节点位置(基于根节点依次递推下来的, 归属于[l,r])
	if r < start || end < l {
		// 更新的前提是 [l,r]理论上应该和[start,end]存在一定程度的重叠, 否则不需要更新
		return
	}

	if start <= l && r <= end {
		// [start, end]包含[l, r], 此时这部分区间的值直接更新即可
		// 这里更新的可能是一个区间, 并不会是一个具体的点
		// 所以被称为: 动态开点
		p := tree[idx]
		p.first += val
		p.second += val
		tree[idx] = p
		return
	}
	// [start, end]和[l, r]存在部分重叠, 需要分割[l,mid]和[mid+1, r], 进行下一部分的更新
	mid := (l + r) >> 1
	// 更新左右子节点
	tree.update(start, end, val, l, mid, 2*idx)
	tree.update(start, end, val, mid+1, r, 2*idx+1)
	// 更新父节点
	p := tree[idx]
	p.first = p.second + max(tree[2*idx].first, tree[2*idx+1].first)
	tree[idx] = p
}

func (tree MyCalendarTwo) Book(start, end int) bool {
	// 尝试预定
	tree.update(start, end-1, 1, 0, 1e9, 1)
	if tree[1].first > 2 {
		tree.update(start, end-1, -1, 0, 1e9, 1)
		return false
	}
	return true
}
