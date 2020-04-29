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

func main() {
	fmt.Println(intToRoman(1994))
}
