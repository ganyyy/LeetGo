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
	strLen, wordsLen, wordLen := len(s), len(words), len(words[0])

	// 这个字典用来记录一次迭代中, 按照wordLen划分的单词的计数.
	// 每增加一个单词, 就在字典中增加一个计数; 每减少一个单词, 就在字典中减少一个计数. 当计数为0时, 就删除这个单词
	// 当字典的长度为0时, 就可以判断这段字母可以由words中的单词组成
	var differWords = make(map[string]int)

	// 找到合适的划分区间

	// 因为在长度为wordLen的word中的任意位置都可以作为起点, 所以需要wordLen次循环
	// 为什么只需要循环wordLen次呢? 因为当循环完一次之后, 就相当于把所有的起点都尝试了一遍, 再来都是重复的
	// 所有的单词都用上的最小长度为wordsLen*wordLen, 所以需要满足 i+wordsLen*wordLen <= strLen
	for i := 0; i < wordLen && i+wordsLen*wordLen <= strLen; i++ {
		clear(differWords)
		// 统计这一段初始字母中, 所有单词的计数
		// 因为words中的每个单词都只能用也必须用一次,
		// 一次可以组成的最长的字符串长度为wordsLen*wordLen
		// 这里相当于截取了 s[i:i+wordsLen*wordLen] , 用来判断是否可以由words中的单词组成
		// 这是起始状态
		differ := differWords
		for j := 0; j < wordsLen; j++ {
			differ[s[i+j*wordLen:i+(j+1)*wordLen]]++
		}
		// 去掉目标单词对应的计数
		for _, word := range words {
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
		}

		// 如果differ中的单词数量为空, 就说明这段字母可以由给定的单词组成
		if len(differ) == 0 {
			ans = append(ans, i)
		}
		// 随后, 开始以步近为wordLen的步长进行迭代, 每次迭代, 右端点加入一个单词, 左端点删除一个单词
		// [i, i+wordsLen*wordLen] -> [i+wordLen, strLen-wordsLen*wordLen+1]
		for start := i + wordLen; start < strLen-wordsLen*wordLen+1; start += wordLen {
			// 右端点加入一个单词
			word := s[start+(wordsLen-1)*wordLen : start+wordsLen*wordLen]
			differ[word]++
			if differ[word] == 0 {
				delete(differ, word)
			}
			// 左端点删除一个单词
			word = s[start-wordLen : start]
			differ[word]--
			if differ[word] == 0 {
				delete(differ, word)
			}
			// 如果differ中的单词数量为空, 就说明这段字母可以由给定的单词组成
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
