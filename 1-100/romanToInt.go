package main

import "fmt"

func romanToInt(s string) int {
	mk := map[uint8]int{
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

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
