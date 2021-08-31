package main

import (
	"strconv"
	"strings"
)

func compareVersionBad(version1 string, version2 string) int {
	var sp1, sp2 = strings.Split(version1, "."), strings.Split(version2, ".")

	// 保证数组1是最小的
	var swap bool
	if len(sp1) > len(sp2) {
		swap = true
		sp1, sp2 = sp2, sp1
	}

	var result = func(a int) int {
		if swap {
			return -a
		}
		return a
	}

	var compare = func(a, b int) int {
		var val int
		if a > b {
			val = 1
		} else {
			val = -1
		}
		return result(val)
	}

	for i := 0; i < len(sp1); i++ {
		var a, _ = strconv.Atoi(sp1[i])
		var b, _ = strconv.Atoi(sp2[i])
		if a != b {
			return compare(a, b)
		}
	}
	for j := len(sp1); j < len(sp2); j++ {
		var v, _ = strconv.Atoi(sp2[j])
		if v != 0 {
			return result(-1)
		}
	}
	return 0
}

func compareVersion(version1, version2 string) int {
	n, m := len(version1), len(version2)
	i, j := 0, 0
	for i < n || j < m {
		x := 0
		for ; i < n && version1[i] != '.'; i++ {
			x = x*10 + int(version1[i]-'0')
		}
		i++ // 跳过点号
		y := 0
		for ; j < m && version2[j] != '.'; j++ {
			y = y*10 + int(version2[j]-'0')
		}
		j++ // 跳过点号
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}
