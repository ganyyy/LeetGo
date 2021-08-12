package main

var base [11]int

func init() {
	base[0] = 0
	base[1] = 1

	var cur = 1

	for i := 2; i <= 10; i++ {
		cur *= 10
		base[i] = base[i-1]*10 + cur
	}
}

func countDigitOne(n int) int {
	var cur = 1
	var cnt = 0
	var tmp = n
	var ret int
	for tmp != 0 {
		var t = tmp % 10
		if t > 1 {
			// 该位置大于1的情况, 需要加上 对应位置的基础个数 * 具体的倍数
			// 再加上该位置1开头的数字个数
			ret += base[cnt]*t + cur
		}
		if t == 1 {
			// 等于1时, 需要加上所有非当前位的数字的个数 + 1(10000...) + 对应位置的基础个数
			ret += n%cur + 1 + base[cnt]
		}
		tmp /= 10
		cnt++
		cur *= 10
	}
	return ret
}
func main() {
	println(countDigitOne(114514))
}
