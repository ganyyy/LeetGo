package main

func platesBetweenCandles(s string, queries [][]int) []int {

	var cnt = make([]int, len(s)) // * 的计数
	var left, right = make([]int, len(s)+1), make([]int, len(s)+1)

	var cur int
	for i, v := range s {
		left[i+1] = left[i]
		if v == '*' {
			cur++
		} else {
			left[i+1] = i
		}
		cnt[i] = cur
	}
	for i := len(s) - 1; i >= 0; i-- {
		right[i] = right[i+1]
		if s[i] == '|' {
			right[i] = i
		}
	}

	left = left[1:]
	right = right[:len(right)-1]

	// fmt.Println(left, right)
	// fmt.Println(cnt)

	var ret = make([]int, len(queries))
	for i, query := range queries {
		if query[0] == query[1] {
			continue
		}
		var l = right[query[0]]
		var r = left[query[1]]
		// fmt.Println(l, r, cnt[l], cnt[r])
		if l >= r {
			continue
		}
		ret[i] = cnt[r] - cnt[l]
	}
	return ret
}
