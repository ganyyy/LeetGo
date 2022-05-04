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

	// 这个可以理解为: 从头开始并不停地扣除首个数字

	// 1,2,3,4 -> 2,4 -> 2
	// 1,2,3,4,5,6 -> 2,4,6 -> 4

	for remain > 1 {
		if leftToRight || remain&1 == 1 {
			// 从左到右, 或者从右向左且剩余数量为奇数时, 才会消除第一个数字
			ret += step
		}
		leftToRight = !leftToRight
		step *= 2
		remain /= 2
	}

	// 最后剩下的一定是经过若干轮的消除后, 剩下的首个数字
	return ret
}

func lastRemaining2(n int) int {
	/*
		i 为数组的长度
		定义f[i] 为 左->右->左执行删除后, 剩余的结果
		   f'[i]为 右->左->右执行删除后, 剩余的结果
		因为对称性, 可以得到
			f[i]+f'[i] = i+1 ①
		进行过一次左->右删除之后
			[1,2,3,4,5,6...i](长度为i) -> [2,4,6,8,...i^1](长度为i/2)
		由关联性可得(序列整体/2的f'[i/2], 计算结果*2得f[i])
			f[i] = f'[i/2]*2 ②
		由①和②可得
			f'[i] = i+1-f[i]
			f[i] = 2*(i/2+1-f[i/2])
		f[1] = 1
	*/
	if n == 1 {
		return 1
	}
	return 2 * (n/2 + 1 - lastRemaining2(n/2))
}
