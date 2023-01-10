package main

import "math"

func crackSafe(n int, k int) string {
	kn := int(math.Pow(float64(k), float64(n)))
	kn1 := kn / k
	num := make([]int, kn1)
	for i := range num {
		num[i] = k - 1
	}
	s := make([]rune, kn+(n-1))
	for i := range s {
		s[i] = '0'
	}
	// 从大到小, 递减迭代
	for i, node := n-1, 0; i < len(s); i++ {
		s[i] = rune(num[node]) + '0'
		num[node]--
		node = node*k - int(s[i-(n-1)]-'0')*kn1 + num[node] + 1
	}
	return string(s)
}
