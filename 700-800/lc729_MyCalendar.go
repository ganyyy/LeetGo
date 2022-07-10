package main

import "sort"

type MyCalendar struct {
	// RB Tree的那个?
	AllPair []CalendarPair
}

type CalendarPair struct {
	Day int
	End int
}

func Constructor3() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) search(start int, end int) bool {
	if len(c.AllPair) == 0 {
		// 首个元素
		c.AllPair = append(c.AllPair, CalendarPair{Day: start, End: end})
		return true
	}
	var idx = sort.Search(len(c.AllPair), func(i int) bool {
		return c.AllPair[i].Day >= start
	})
	// 首个小于start的最大值
	var isFirst = idx == 0
	if !isFirst {
		idx--
	}
	var cur = c.AllPair[idx]
	// defer func() {
	//     fmt.Println(idx, start, end,c.AllPair, cur)
	// }()

	if (cur.Day <= start && start < cur.End) || (start <= cur.Day && cur.Day < end) {
		// 存在重叠
		return false
	}
	if idx+1 < len(c.AllPair) {
		// 看后继
		var next = c.AllPair[idx+1]
		if next.Day < end {
			return false
		}
	}
	// idx是要插入的位置
	if !isFirst {
		idx++
	}
	c.AllPair = append(c.AllPair, CalendarPair{})
	copy(c.AllPair[idx+1:], c.AllPair[idx:])
	c.AllPair[idx] = CalendarPair{Day: start, End: end}
	return true
}

func (this *MyCalendar) Book(start int, end int) bool {
	return this.search(start, end)
}

type MyCalendar2 struct {
	tree, lazy map[int]bool
}

func Constructor4() MyCalendar2 {
	return MyCalendar2{map[int]bool{}, map[int]bool{}}
}

func (c MyCalendar2) query(start, end, l, r, idx int) bool {
	if r < start || end < l {
		return false
	}
	if c.lazy[idx] { // 如果该区间已被预订，则直接返回
		return true
	}
	if start <= l && r <= end {
		return c.tree[idx]
	}
	mid := (l + r) >> 1
	return c.query(start, end, l, mid, 2*idx) ||
		c.query(start, end, mid+1, r, 2*idx+1)
}

func (c MyCalendar2) update(start, end, l, r, idx int) {
	if r < start || end < l {
		return
	}
	if start <= l && r <= end {
		c.tree[idx] = true
		c.lazy[idx] = true
	} else {
		mid := (l + r) >> 1
		// 左孩子
		c.update(start, end, l, mid, 2*idx)
		// 右孩子
		c.update(start, end, mid+1, r, 2*idx+1)
		c.tree[idx] = true
		if c.lazy[2*idx] && c.lazy[2*idx+1] {
			c.lazy[idx] = true
		}
	}
}

func (c MyCalendar2) Book(start, end int) bool {
	// 1为树节点的根,
	if c.query(start, end-1, 0, 1e9, 1) {
		return false
	}
	c.update(start, end-1, 0, 1e9, 1)
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
