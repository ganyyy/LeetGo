package main

func addDigits(num int) int {

	// 拿三位数 abc 举例
	// a*100+b*10+c ①
	// a+b+c        ②
	// ①-② 得 99*a+9*b ③
	// ③ ÷ 9 == 0
	// 每一次缩小的值都是9的倍数

	if num < 9 {
		return num
	}
	if v := num % 9; v == 0 {
		return 9 // 和9取余的结果为0, 那么最终相加的结果一定是9
	} else {
		return v
	}
}
