package main

import "strconv"

func simplifiedFractions(n int) []string {

	var tmp = make([]string, 0, n)

	for j := 1; j < n; j++ {
		for i := j + 1; i <= n; i++ {
			if check(j, i) {
				continue
			}
			tmp = append(tmp, strconv.Itoa(j)+"/"+strconv.Itoa(i))
		}
	}
	return tmp
}

func check(a, b int) bool {
	// 需要判断一下是否存在非1的公约数
	// a小b大
	for a != 0 {
		a, b = b%a, a
	}
	return b != 1
}
