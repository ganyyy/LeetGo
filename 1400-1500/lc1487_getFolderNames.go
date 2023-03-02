package main

import "strconv"

func getFolderNames(names []string) []string {
	ans := make([]string, len(names))
	index := map[string]int{}
	for p, name := range names {
		// 第一步: 找原始的名字. name1, name1(1)
		i := index[name]
		if i == 0 {
			index[name] = 1
			ans[p] = name
			continue
		}
		// 这一步可以保证: 即使按照 a, a(1), a 的顺序出现, 最终的结果也会是 a,a(1),a(2)
		for index[name+"("+strconv.Itoa(i)+")"] > 0 {
			i++
		}
		// 第二步, 把标记的名字也带上. name1(1), name(1)(2)
		t := name + "(" + strconv.Itoa(i) + ")"
		ans[p] = t
		index[name] = i + 1
		index[t] = 1
	}
	return ans
}
