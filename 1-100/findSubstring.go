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

func findSubstring2(s string, words []string) (ans []int) {
	// 总长度ls
	// m个单词
	// 每个单词的长度为n
	ls, m, n := len(s), len(words), len(words[0])

	// 找到合适的划分区间
	// 因为单词的长度为n, 所以n次为一轮
	// 所有的单词都用上的最小长度为m*n, 所以需要满足 i+m*n <= ls
	for i := 0; i < n && i+m*n <= ls; i++ {
		// 统计这一段初始字母中, 所有单词的计数
		differ := map[string]int{}
		for j := 0; j < m; j++ {
			differ[s[i+j*n:i+(j+1)*n]]++
		}
		// 去掉目标单词对应的计数
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}
		for start := i; start < ls-m*n+1; start += n {
			if start != i {
				// 右端点加入一个单词
				word := s[start+(m-1)*n : start+m*n]
				differ[word]++
				if differ[word] == 0 {
					delete(differ, word)
				}
				// 左端点删除一个单词
				word = s[start-n : start]
				differ[word]--
				if differ[word] == 0 {
					delete(differ, word)
				}
			}
			// 如果differ中的单词数量为空, 就说明这段字母由给定的单词组成
			if len(differ) == 0 {
				ans = append(ans, start)
			}
		}
	}
	return
}

func main() {
	/**
	  "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	  ["fooo","barr","wing","ding","wing"]
	*/
	fmt.Println(findSubstring2("lingmindraboofooowingdingbarrwingmonkeypoundcake", []string{"fooo", "barr", "wing", "ding", "wing"}))
}
