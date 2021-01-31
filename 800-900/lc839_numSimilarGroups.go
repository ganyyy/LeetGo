package main

import "fmt"

func numSimilarGroups(strs []string) int {
	if len(strs) == 0 {
		return 0
	}

	// 并查集的数量
	// 啊这, 模板我都会背了

	// 保存一下根节点的所有异位词
	var m = make(map[string]map[string]bool)
	// 初始化一下. 当前字符串到父字符串的映射
	var fa = make(map[string]string, len(strs))
	for _, s := range strs {
		// 字符串->父字符串
		fa[s] = s
		// 该节点含有的子节点
		m[s] = map[string]bool{s: true}
	}

	var find func(s string) string
	find = func(s string) string {
		if s != fa[s] {
			fa[s] = find(fa[s])
		}
		return fa[s]
	}

	// 将两个字符串合并到一起
	var merge = func(sf, st string) {
		sf, st = find(sf), find(st)
		if sf == st {
			return
		}
		if st < sf {
			st, sf = sf, st
		}
		fa[sf] = st
		for s := range m[sf] {
			m[st][s] = true
		}
		delete(m, sf)
		return
	}

	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			var s1, s2 = find(strs[i]), find(strs[j])
			if s1 == s2 {
				continue
			}
			var ms1, ms2 = m[s1], m[s2]
			if ms1 == nil || ms2 == nil {
				continue
			}
		end:
			for k1 := range ms1 {
				for k2 := range ms2 {
					if isSame(k1, k2) {
						merge(s1, s2)
						break end
					}
				}
			}
		}
	}
	return len(m)
}
func numSimilarGroups2(strs []string) int {
	var ln = len(strs)
	if ln == 0 {
		return 0
	}

	// 并查集的数量
	// 啊这, 模板我都会背了

	// 去重
	var m = make(map[string]bool, ln)
	for _, s := range strs {
		m[s] = true
	}

	if len(m) <= 1 {
		return len(m)
	}

	// 重组
	strs = make([]string, len(m))
	var i int
	for s := range m {
		strs[i], i = s, i+1
	}
	ln = len(strs)
	// 初始化一下. 当前字符串到父字符串的映射
	var fa = make([]int, ln)
	for i := range strs {
		// 这里主要是保证没有发生过交换
		fa[i] = -1
	}

	var find = func(i int) int {
		for {
			r := fa[i]
			if r == i || r == -1 {
				return i
			}
			i = r
		}
	}

	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			var s1, s2 = find(i), find(j)
			if s1 == s2 {
				continue
			}
			if isSame(strs[s1], strs[s2]) {
				// 向小的地方走, 因为会先便利到前边的
				fa[s1] = s2
			}
		}
	}
	fmt.Print(fa, strs)
	var res int
	for _, v := range fa {
		if v == -1 {
			res++
		}
	}
	return res
}

//isSame 判断两个字符串是不是异位词
func isSame(s1, s2 string) bool {
	var cnt int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			cnt++
			if cnt > 2 {
				return false
			}
		}
	}
	return true
}

func numSimilarGroups3(strs []string) int {
	var ln = len(strs)
	if ln == 0 {
		return 0
	}

	// 初始化一下. 当前字符串到父字符串的映射
	var fa = make([]int, ln)
	for i := range strs {
		fa[i] = i
	}

	var find func(i int) int
	find = func(i int) int {
		if i != fa[i] {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}

	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			if isSame(strs[i], strs[j]) {
				fa[find(i)] = find(j)
			}
		}
	}
	var res int
	for i := range fa {
		if find(i) == i {
			res++
		}
	}
	return res
}

func main() {
	var strs = []string{"nmiwx", "mniwx", "wminx", "mnixw", "xnmwi"}

	for i := 0; i < 10; i++ {
		fmt.Println("\t", numSimilarGroups2(strs))
	}
}
