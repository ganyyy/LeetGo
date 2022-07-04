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

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
