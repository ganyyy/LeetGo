package main

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {

	// 取中间的两个x和两个y, 计算面积
	// 需要确定是否相交欸..

	var a, b = (ax2 - ax1) * (ay2 - ay1), (bx2 - bx1) * (by2 - by1)
	if (ax2 <= bx1) || (ax1 >= bx2) || (ay1 >= by2) || (ay2 <= by1) {
		// 不相交
		return a + b
	}

	// 如果相交, 去掉公共部分
	return a + b - (min(ax2, bx2)-max(ax1, bx1))*(min(ay2, by2)-max(ay1, by1))
}
