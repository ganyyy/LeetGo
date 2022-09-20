package main

func kSimilarity(s1, s2 string) (step int) {
	type pair struct {
		s string // 交换后的字符串
		i int    // 以及使用该字符应该从哪个位置开始迭代
	}
	q := []pair{{s: s1}}
	vis := map[string]bool{s1: true} // 记忆化搜索
	for n := len(s1); ; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			s, i := p.s, p.i
			if s == s2 {
				return
			}
			// i: s[i] != s2[i]
			for i < n && s[i] == s2[i] {
				i++
			}
			t := []byte(s)
			// j: s[j] == s2[i], 不过如果s[j] == s2[j], 那么也不应该进行交换!
			for j := i + 1; j < n; j++ {
				if s[j] == s2[i] && s[j] != s2[j] {
					t[i], t[j] = t[j], t[i]
					if t := string(t); !vis[t] {
						vis[t] = true
						q = append(q, pair{t, i + 1})
					}
					t[i], t[j] = t[j], t[i]
				}
			}
		}
	}
}
