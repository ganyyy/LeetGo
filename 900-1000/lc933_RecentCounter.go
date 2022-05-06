//go:build ignore

package main

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (c *RecentCounter) Ping(t int) int {
	var queue = c.queue
	var i int
	for ; i < len(queue); i++ {
		if queue[i] >= t-3000 {
			break
		}
	}
	queue = queue[i:]
	queue = append(queue, t)
	c.queue = queue
	return len(queue)
}
