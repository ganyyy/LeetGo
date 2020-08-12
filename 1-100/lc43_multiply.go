package main

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1, l2 := len(num1), len(num2)
	res := make([]int, l1+l2)

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			mul := int((num1[i] - '0') * (num2[j] - '0'))
			res[i+j+1] += mul % 10
			res[i+j] += mul / 10
		}
	}
	ret := make([]byte, len(res))
	fmt.Println(res)
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] >= 10 {
			res[i-1] += res[i] / 10
			res[i] %= 10
		}
		ret[i] = byte(res[i]) + '0'
	}
	if ret[0] == '0' {
		ret = ret[1:]
	}
	return string(ret)
}
func multiply2(num1 string, num2 string) string {
	res := make([]int, len(num1)+len(num2))

	// 倒序相乘
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			m := int(num1[i]-'0') * int(num2[j]-'0')
			// 加进位
			m += res[i+j+1]
			// 高位
			res[i+j] += m / 10
			// 低位
			res[i+j+1] = m % 10
		}
	}

	// 去零
	var i int
	// 至少保留一位, 允许是 0
	for ; i < len(res)-1 && res[i] == 0; i++ {
	}

	// 填充结果
	r := len(res) - i
	s := make([]byte, r)
	for t := 0; t < r; t++ {
		s[t] = byte(res[t+i] + '0')
	}
	return string(s)
}

func main() {
	fmt.Println(multiply("123", "456"))
}
