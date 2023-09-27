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

func addBinary2(a string, b string) string {
	// b始终最长
	if len(a) > len(b) {
		a, b = b, a
	}
	res := []byte(b)
	sub := len(b) - len(a)
	// 进位标志
	var add byte
	for i := len(b) - 1; i >= 0; i-- {
		cur := res[i] - '0' + add
		if i-sub >= 0 {
			cur += a[i-sub] - '0'
		}
		if cur >= 2 {
			res[i] = cur - 2 + '0'
			add = 1
		} else {
			res[i] = cur + '0'
			add = 0
		}
		if add == 0 && i-sub < 0 {
			break
		}
	}
	if add == 1 {
		return "1" + string(res)
	} else {
		return string(res)
	}
}

func addBinary3(a string, b string) string {
	// b始终最长
	if len(a) > len(b) {
		a, b = b, a
	}
	res := make([]byte, 1+len(b))
	copy(res[1:], b)
	ia := len(a) - 1
	// 进位标志
	var add byte
	for ib := len(b) - 1; ib >= 0; ib-- {
		cur := b[ib] + add
		if ia >= 0 {
			cur += a[ia] - '0'
			ia--
		}
		if cur >= '2' {
			cur -= 2
			add = 1
		} else {
			add = 0
		}
		res[ib+1] = cur
		if add == 0 && ia < 0 {
			break
		}
	}
	if add == 1 {
		res[0] = '1'
	} else {
		res = res[1:]
	}
	return string(res)
}

func main() {
	fmt.Println(addBinary2(
		"1011",
		"110"),
	)
}
