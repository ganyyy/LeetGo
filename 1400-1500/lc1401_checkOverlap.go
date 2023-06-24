//go:build ignore

package main

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	dist := 0
	// .....x1.....x2.....
	//   c      c      c
	// 在两边的时候, 需要计算一下其距离的平方
	if xCenter < x1 || xCenter > x2 {
		dist += min((x1-xCenter)*(x1-xCenter), (x2-xCenter)*(x2-xCenter))
	}
	// 和x同理, 计算y距离的平方
	if yCenter < y1 || yCenter > y2 {
		dist += min((y1-yCenter)*(y1-yCenter), (y2-yCenter)*(y2-yCenter))
	}

	/*
	   将矩形对应的两个横坐标, 两个纵坐标无限延长, 将整个二维区间分为9块区域
	   然后基于 圆心所处的位置分开进行讨论
	   当圆心处于外围的8个区域时, 需要分别计算一下圆心矩形之间的距离的平方
	   处于斜对角的4个区域时, 圆心到矩形的最短距离就是到离他最近的端点的距离
	   处于正方向的4个区域时, 圆心到矩形的最短距离就是到离他最近的边的距离(比如正上就是min(y1-yCenter, y2-yCenter))
	*/

	return dist <= radius*radius
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
