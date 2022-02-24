package main

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(num1 string, num2 string) string {

	var parse = func(s string) (int, int) {
		var sp1 = strings.Split(s, "+")
		var n1, _ = strconv.Atoi(sp1[0])
		var n2, _ = strconv.Atoi(sp1[1][:len(sp1[1])-1])
		return n1, n2
	}

	var a1, a2 = parse(num1)
	var b1, b2 = parse(num2)

	return strconv.Itoa(a1*b1-a2*b2) + "+" + strconv.Itoa(a1*b2+a2*b1) + "i"
}
