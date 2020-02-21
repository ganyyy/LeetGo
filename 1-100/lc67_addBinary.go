package main

import "fmt"

func addBinary(a string, b string) string {
	i := 0
	for ; i < len(a); i++ {
		if a[i] != '0' {
			break
		}
	}
	a = a[i:]
	for i = 0; i < len(b); i++ {
		if b[i] != '0' {
			break
		}
	}
	b = b[i:]
	if len(a) > len(b) {
		a, b = b, a
	}
	la, lb := len(a), len(b)
	if la == lb && lb == 0 {
		return "0"
	}
	res := make([]byte, lb+1)
	res[lb] = '0'
	i, add := lb-1, byte(0)
	sub := lb - la
	for ; i >= sub; i-- {
		res[i+1] = a[i-sub] + b[i] - '0' + add
		if res[i+1] >= '2' {
			add = 1
			res[i+1] -= 2
		} else {
			add = 0
		}
	}
	for ; i >= 0; i-- {
		res[i+1] = b[i] + add
		if res[i+1] >= '2' {
			res[i+1] -= 2
			add = 1
		} else {
			add = 0
		}
	}
	if add == 0 {
		res = res[1:]
	} else {
		res[0] = '1'
	}
	return string(res)
}

func main() {
	fmt.Println(addBinary("1111110", "0"))
}
