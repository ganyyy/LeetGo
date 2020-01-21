package main

import "fmt"

func findSubstring(s string, words []string) []int {
	var res []int
	ls, lw := len(s), len(words)
	if ls == 0 {
		return res
	}
	if lw == 0 {
		return res
	}
	wl := len(words[0])
	ml := lw * wl
	if ls < ml {
		return res
	}
	m1 := make(map[string]int, lw)
	for _, v := range words {
		m1[v] += 1
	}
	for i := 0; i < wl; i++ {
		res = append(res, sub(i, m1, s[i:], ls, lw, wl, ml)...)
	}
	return res
}

func sub(left int, m1 map[string]int, s string, ls, lw, wl, ml int) []int {
	var res []int
	for i := 0; i <= ls-ml; i += wl {
		m := make(map[string]int)
		if i+ml > len(s) {
			break
		}
		ts := s[i : i+ml]
		j := 0
		for ; j < lw; j++ {
			start := j * wl
			end := start + wl
			ss := ts[start:end]
			if v, ok := m1[ss]; ok {
				if mv, ok := m[ss]; ok {
					mv++
					// 数量过了直接跳出
					if mv > v {
						break
					} else {
						m[ss] = mv
					}
				} else {
					m[ss] = 1
				}
			} else {
				// 找不到直接跳出
				break
			}
		}
		// 相等就插进去
		if j == lw {
			res = append(res, i+left)
		}
	}
	return res
}

func main() {
	/**
	"lingmindraboofooowingdingbarrwingmonkeypoundcake"
	["fooo","barr","wing","ding","wing"]
	*/
	fmt.Println(findSubstring("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
}
