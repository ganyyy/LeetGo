package main

import "math"

func shortestToChar(s string, c byte) []int {
	var ret = make([]int, len(s))

	var queue []int
	for i := range s {
		if s[i] == c {
			ret[i] = 0
			queue = append(queue, i)
		} else {
			ret[i] = math.MaxInt32
		}
	}
	var nextQueue []int
	for len(queue) != 0 {
		var ln = len(queue)
		for i := 0; i < ln; i++ {
			var cur = queue[i]
			var pre, next = cur - 1, cur + 1
			if pre >= 0 && ret[pre] == math.MaxInt32 {
				ret[pre] = ret[cur] + 1
				nextQueue = append(nextQueue, pre)
			}
			if next < len(ret) && ret[next] == math.MaxInt32 {
				ret[next] = ret[cur] + 1
				nextQueue = append(nextQueue, next)
			}
		}
		queue, nextQueue = nextQueue, queue
		nextQueue = nextQueue[:0]
	}

	return ret
}
