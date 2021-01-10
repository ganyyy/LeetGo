package main

import (
	"sort"
	"unsafe"
)

func smallestStringWithSwaps(s string, pairs [][]int) string {
	// 通过 所有可以转换的节点, 找到所有的转换路径.
	// 然后分批进行转换, 求对应的最小字典序 组合后返回

	var f = make([]int, len(s))
	for i := range f {
		f[i] = i
	}

	var find func(a int) int

	find = func(a int) int {
		if f[a] != a {
			f[a] = find(f[a])
		}
		return f[a]
	}

	var merge = func(from, to int) {
		if from == to {
			return
		}
		f[find(from)] = find(to)
	}

	for _, p := range pairs {
		merge(p[0], p[1])
	}

	var m = make(map[int][]int, len(s))

	// 挑选
	for i := range f {
		// 这里重新查找一下父节点, 就可以避免合并不充分的情况
		var p = find(f[i])
		m[p] = append(m[p], i)
	}

	var res = make([]byte, len(s))
	// 排序
	for _, bs := range m {
		var t = make([]byte, len(bs))
		for i, v := range bs {
			t[i] = s[v]
		}
		sort.Slice(t, func(a, b int) bool {
			return t[a] < t[b]
		})
		for i, v := range bs {
			res[v] = t[i]
		}
	}
	return *(*string)(unsafe.Pointer(&res))
}
