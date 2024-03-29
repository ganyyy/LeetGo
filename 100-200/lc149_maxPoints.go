package main

import "unsafe"

func maxPointsOld(points [][]int) int {
	// 叉乘的方式计算, 斜率可能会出现 除数为0的情况

	// 每一个点都需要给前边所有的点计算对应的结果

	// 暴力求解

	if len(points) < 3 {
		return len(points)
	}

	var ret int

	for i := 0; i < len(points); i++ {
		var same = 1

		for j := i + 1; j < len(points); j++ {
			// cnt 表示同一直线上的点
			var cnt int

			var p1, p2 = points[i], points[j]

			if p1[0] == p2[0] && p1[1] == p2[1] {
				same++
			} else {
				cnt++

				var xDiff, yDiff = p1[0] - p2[0], p1[1] - p2[1]

				for k := j + 1; k < len(points); k++ {
					var p3 = points[k]
					// 避免除数为0的情况: 改成乘法
					// 这个转换公式就是: (y2-y1)/(x2-x1) = (y3-y1)/(x3-x1) => (y2-y1)*(x3-x1) = (y3-y1)*(x2-x1)
					if xDiff*(p1[1]-p3[1]) == yDiff*(p1[0]-p3[0]) {
						// 如果在同一条线上...
						cnt++
					}
				}
			}

			if same+cnt > ret {
				ret = same + cnt
			}
		}
	}

	return ret
}

func maxPoints(points [][]int) int {
	count := len(points)
	if count < 3 {
		// 优化: 点数少于3个直接返回
		return count
	}

	var ret int
	var slope = make(map[int]int)

	for i, point := range points {
		if ret >= count-i || ret >= count/2 {
			// 优化: 已经获取的最多的同直线点数 > 剩余点数
			// 或者前边某次迭代获取到的同直线点数已经超过了一半, 都没必要继续迭代了
			// i 越小, 可能同直线的点数就会越多. 往后i越大, 剩余可选的点数也就越少, 必然不能超过ret
			break
		}
		clear(slope)
		for _, nextPoint := range points[i+1:] {
			deltaX := point[0] - nextPoint[0]
			deltaY := point[1] - nextPoint[1]
			if deltaX == 0 {
				// 竖直线
				deltaY = 1
			} else if deltaY == 0 {
				// 水平线
				deltaX = 1
			} else {
				if deltaY < 0 {
					// 用于归一化: 将斜率转换成一个具体的数字
					// 可以这么理解: 协议为 Δy/Δx, Δy或者Δx小于0时, 斜率应该是一致的
					// 同样的, 同时小于0和同时大于0的情况下, 斜率也是一样的
					// 这里针对的是单方面小于0的情况.
					deltaX, deltaY = -deltaX, -deltaY
				}
				// 同时除以最大公约数, 保证斜率是一个最简分数
				xy := gcd(abs149(deltaX), abs149(deltaY))
				deltaX /= xy
				deltaY /= xy
			}
			// 将deltaX和deltaY组合成一个数字, 用于标记斜率
			var sp = *(*int)(unsafe.Pointer(&[2]int32{int32(deltaX), int32(deltaY)}))
			// 统计同一斜率的点数, 并更新最大值
			cnt := slope[sp] + 1
			slope[sp] = cnt
			if ret < cnt {
				ret = cnt
			}
		}
	}
	// +1是算上points[i]那个点位自身的值
	return ret + 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs149(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
