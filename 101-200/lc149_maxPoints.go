package main

func maxPoints(points [][]int) int {
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
