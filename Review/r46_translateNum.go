package main

// 青蛙跳台阶系列 翻版...
func translateNum(num int) int {
	if num < 9 {
		return 1
	}
	// 小于10或者大于25都只有一种情况
	// 如 01 只能是ab, 26只能是cg
	if ba := num % 100; ba < 10 || ba > 25 {
		return translateNum(num / 10)
	}
	// 一种+两种的集合
	return translateNum(num/10) + translateNum(num/100)
}

func main() {

}
