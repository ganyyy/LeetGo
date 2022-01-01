package main

func lastRemainingOvertime(n int) int {
	// 搞个铁定会超时的写法

	var tmp = make([]int, n)
	for i := range tmp {
		tmp[i] = i + 1
	}

	var leftToRight = true
	for len(tmp) > 1 {
		var ln = len(tmp)
		if leftToRight {
			var start int
			for i := 0; i < ln; i++ {
				if i&1 == 1 {
					tmp[start] = tmp[i]
					start++
				}
			}
			tmp = tmp[:start]
		} else {
			var start = len(tmp)
			for i := 0; i < ln; i++ {
				if i&1 == 1 {
					start--
					tmp[start] = tmp[ln-1-i]
				}
			}
			tmp = tmp[start:]
		}
		leftToRight = !leftToRight
	}

	return tmp[0]
}

func lastRemaining(n int) int {
	var remain = n
	var leftToRight = true
	var ret = 1
	var step = 1

	for remain > 1 {
		if leftToRight || remain&1 == 1 {
			// 从左到右, 或者从右向左且数组为奇数时, 才会消除数组的第一个数字
			ret += step
		}
		leftToRight = !leftToRight
		step *= 2
		remain /= 2
	}

	// 最后剩下的一定是经过若干轮的消除后, 剩下的首个数字
	return ret
}
