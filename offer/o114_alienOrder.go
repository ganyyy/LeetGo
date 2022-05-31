package main

func alienOrder(words []string) string {
	g := map[byte][]byte{}
	inDeg := map[byte]int{}
	for _, c := range words[0] {
		inDeg[byte(c)] = 0
	}
next:
	for i := 1; i < len(words); i++ {
		s, t := words[i-1], words[i]
		for _, c := range t {
			inDeg[byte(c)] += 0
		}
		for j := 0; j < len(s) && j < len(t); j++ {
			if s[j] != t[j] {
				g[s[j]] = append(g[s[j]], t[j])
				inDeg[t[j]]++
				continue next
			}
		}
		// 此时已经不满足条件了, 直接返回空串即可
		if len(s) > len(t) {
			return ""
		}
	}

	order := make([]byte, len(inDeg))
	q := order[:0]
	for u, d := range inDeg {
		if d == 0 {
			q = append(q, u)
		}
	}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		for _, v := range g[u] {
			if inDeg[v]--; inDeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	if cap(q) == 0 {
		return string(order)
	}
	return ""
}
