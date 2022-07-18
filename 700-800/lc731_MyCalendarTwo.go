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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
