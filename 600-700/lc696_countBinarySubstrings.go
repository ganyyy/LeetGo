package main

import "fmt"

func countBinarySubstrings(s string) int {
	var cnt [2]int
	var res int

	var check = func(i int, v byte) int {
		cnt[v-'0']++
		if i == 0 {
			return 0
		}
		var opt byte = '0'
		if v == '0' {
			opt = '1'
		}
		// 这是断了
		if opt == s[i-1] {
			return 1
		}
		// 看看最多能合出来几个
		var r int
		for index := 0; index < min(cnt[1], cnt[0]); index++ {
			if s[i-index] == v {
				r++
			} else {
				break
			}
		}
		var c int
		for ; c < r; c++ {
			if s[i-r-c] != opt {
				break
			}
		}
		if c == r && r != 0 {
			return 1
		}
		return 0
	}

	for i, v := range s {
		res += check(i, byte(v))
	}

	return res
}

func countBinarySubstrings2(s string) int {
	// last 前一个字符的数量
	// cur 当前字符的数量
	// res 最终的结果
	var last, cur, res int
	// 当前字符的数量默认为1
	cur = 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			// 如果相对, 当前计数+1,
			cur++
		} else {
			// 否则转换当前计数为上一个的计数, 并且重置当前的计数
			last = cur
			cur = 1
		}
		// 如果前一个字符的数量大于等于当前字符的数量,
		// 说明一定存在一种排列满足 题目要求
		if last >= cur {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println("res is ", countBinarySubstrings("10101"))
}
