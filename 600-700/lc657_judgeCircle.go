package main

func judgeCircle(moves string) bool {
	// 搞两个栈, 一个存储上下移动
	// 一个存储左右移动

	// 直接对比数量, 不是更好么?
	var l, r, u, d int
	for _, v := range moves {
		switch byte(v) {
		case 'L':
			l++
		case 'R':
			r++
		case 'U':
			u++
		case 'D':
			d++
		}
	}
	return l == r && u == d
}
