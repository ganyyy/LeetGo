package main

import "fmt"

func isHappy(n int) bool {
	var t, i int
	m := make(map[int]bool)
	for n != 1 {
		for t = 0; n != 0; n /= 10 {
			i = n % 10
			t += i * i
		}
		if _, ok := m[t]; ok {
			return false
		}
		m[t] = true
		n = t
		// 如何判断是不是进入循环了呢?
	}
	return true
}

func main() {
	fmt.Println(isHappy(19))
}
