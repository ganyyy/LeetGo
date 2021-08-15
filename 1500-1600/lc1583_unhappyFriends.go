package main

import "math"

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	// 配对关系映射
	var p = make([]int, n)
	for _, v := range pairs {
		p[v[0]] = v[1]
		p[v[1]] = v[0]
	}
	var ret int

	// i    是第几个人
	// pair 是i对应的配对人员
	for i, pair := range p {
		// pr是和i的亲密度高于 当前配对人  的其他人
		for _, pr := range preferences[i] {
			if pr == pair {
				// 如果到了当前配对的人, 没必要继续查下去了.
				// 因为往后的亲密度都低于当前配对的人
				break
			}
			// ppair 是 pr的配对人
			var ppair = p[pr]

			// ppn  是pr和当前配对人的亲密度
			// pin  是pr和i的亲密度
			var pin, ppn = -1, -1
			for pan, par := range preferences[pr] {
				// 还是有优化的空间. 比如, 如果提前找到了i, 就没必要继续找下去了
				if par == i {
					pin = pan
					if ppn == -1 {
						ppn = math.MaxInt32
					}
					break
				}
				if par == ppair {
					ppn = pan
					if pin != -1 {
						break
					}
					continue
				}

			}
			// 如果 pair 对 i的亲密度高于 pair对当前配对人员的亲密度, 此时 i就是不开心的
			if ppn > pin {
				// 找到一个就够了, 直接跳出
				ret++
				break
			}
		}
	}
	return ret
}
