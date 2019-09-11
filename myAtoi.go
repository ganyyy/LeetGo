package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	res := 0
	flag, start := false, false
	bFlag := false
	for pos := 0; pos < len(str); pos++ {
		c := str[pos]
		switch c {
		case ' ':
			if !start {
				continue
			} else {
				bFlag = true
			}
		case '-':
			if start {
				bFlag = true
			} else {
				start = true
				flag = true
			}
		case '+':
			if start {
				bFlag = true
			} else {
				start = true
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			start = true
			res = res * 10 + int(c) - '0'
			if flag {
				if -res <= math.MinInt32 {
					return math.MinInt32
				}
			} else {
				if res >= math.MaxInt32 {
					return math.MaxInt32
				}
			}
		default:
			bFlag = true
		}
		if bFlag {
			break
		}
	}

	if flag {
		if -res < math.MinInt32 {
			return math.MinInt32
		} else {
			return -res
		}
	} else {
		if res > math.MaxInt32 {
			return math.MaxInt32
		} else {
			return res
		}
	}
}

func main() {
	fmt.Println(myAtoi("-2147483647"))
}
