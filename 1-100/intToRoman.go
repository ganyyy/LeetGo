package main

import (
	"fmt"
	"strings"
)

func intToRoman(num int) string {
	k := []int{
		1, 4, 5, 9,
		10, 40, 50, 90,
		100, 400, 500, 900,
		1000,
	}
	v := []string{
		"I", "IV", "V", "IX",
		"X", "XL", "L", "XC",
		"C", "CD", "D", "CM",
		"M",
	}
	res := strings.Builder{}

	i := len(k) - 1
	for num != 0 {
		tK := k[i]
		if num >= tK {
			d := num / tK
			tV := v[i]
			for j := 0; j < d; j++ {
				res.WriteString(tV)
			}
			num = num % tK
		} else {
			i--
		}
	}
	return res.String()
}

var m = [...]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

var s = [...]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

func intToRoman2(num int) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		if num >= s[i] {
			n := num / s[i]
			var str = m[i]
			// fmt.Println(num, n, s[i], str)
			for i := 0; i < n; i++ {
				sb.WriteString(str)
			}
			num -= n * s[i]
			if num == 0 {
				break
			}
		}
	}

	return sb.String()
}

func main() {
	fmt.Println(intToRoman(1994))
}
