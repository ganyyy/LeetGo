//go:build ignore

package main

import "math"

func minimumJumps(forbidden []int, a int, b int, x int) int {
	if x == 0 {
		return 0
	}
	// 计算边界范围
	var mf int
	for _, f := range forbidden {
		mf = max(mf, f)
	}
	// 证明过程: 略...
	bound := max(mf+a+b, x+b)

	// 题目给出的上限
	var ban = make([]bool, bound+1)
	for _, f := range forbidden {
		ban[f] = true
	}
	// [0]: 往前跳到指定位置的步数
	// [1]: 往后跳到指定位置的步数
	var dist = make([][2]int, bound+1)
	for i := range dist {
		dist[i] = [2]int{math.MaxInt32, math.MaxInt32}
	}
	dist[0][0] = 0

	// [0]: 位置
	// [1]: 方向. 0正向, 1反向
	var queue, next [][2]int
	// 起始位置: 0, 前向
	queue = append(queue, [2]int{0, 0})

	for len(queue) != 0 {
		for _, node := range queue {
			i, dir := node[0], node[1]
			nxtStep := dist[i][dir] + 1
			// 上次是正向的才能往后退
			if dir == 0 {
				back := i - b
				if back == x {
					return nxtStep
				}
				// 满足条件, 并且步数最优(?)
				// 不过, 首次到达的一定是最优解啊...?
				if back >= 0 && !ban[back] && nxtStep < dist[back][1] {
					dist[back][1] = nxtStep
					next = append(next, [2]int{back, 1})
				}
			}
			forward := i + a
			if forward == x {
				return nxtStep
			}
			if forward <= bound && !ban[forward] && nxtStep < dist[forward][0] {
				dist[forward][0] = nxtStep
				next = append(next, [2]int{forward, 0})
			}
		}
		queue, next = next, queue[:0]
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
