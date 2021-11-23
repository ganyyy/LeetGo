package main

import "strings"

var numStr = [10]string{
	"zero",
	"two",
	"six",
	"eight",
	"four",
	"one",
	"five",
	"seven",
	"three",
	"nine",
}

var numChar = [10]byte{
	'z',
	'w',
	'x',
	'g',
	'u',
	'o',
	'f',
	'v',
	't',
	'i',
}

var numInt = [10]int{
	0,
	2,
	6,
	8,
	4,
	1,
	5,
	7,
	3,
	9,
}

func originalDigits(s string) string {
	var set [26]int

	for _, c := range s {
		set[int(c-'a')]++
	}

	var intCnt [10]int
	var total int
	for i, str := range numStr {
		var cnt = set[int(numChar[i]-'a')]
		intCnt[numInt[i]] += cnt
		total += cnt
		for _, c := range str {
			set[int(c-'a')] -= cnt
		}
	}

	var sb strings.Builder

	sb.Grow(total)

	for i, n := range intCnt {
		for j := 0; j < n; j++ {
			sb.WriteByte(byte(i + '0'))
		}
	}

	return sb.String()
}

func main() {
	println(originalDigits("fviefuro"))
}
