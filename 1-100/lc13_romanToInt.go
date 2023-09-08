package main

import "fmt"

func romanToInt(s string) int {
	mk := [...]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	sLen := len(s)
	for index := 0; index < sLen; {
		v := mk[s[index]]
		if index+1 < sLen {
			v1 := mk[s[index+1]]
			// 理论上大的在前, 如果出现小的在前, 那么就说明这个需要-
			if v1 > v {
				res = res - v
				v = v1
				index++
			}
		}
		res += v
		index++
	}
	return res
}

var m3 = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt2(s string) int {
	// 想啥来啥

	// 如果倒序看, 大的直接加, 小的直接减去
	// mark
	var ret = m3[s[len(s)-1]]
	var before = ret
	for i := len(s) - 2; i >= 0; i-- {
		var v = m3[s[i]]
		if v >= before {
			ret += v
		} else {
			ret -= v
		}
		before = v
	}
	return ret
}

func main() {
	fmt.Println(romanToInt2("MCMXCIV"))
}
