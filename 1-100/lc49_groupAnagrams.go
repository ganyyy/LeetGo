package main

import (
	"sort"
	"unsafe"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)

	for _, str := range strs {
		var letters [26]int
		for i := 0; i < len(str); i++ {
			letters[str[i]-'a']++
		}
		if ss, ok := m[letters]; !ok {
			ss := make([]string, 0)
			m[letters] = append(ss, str)
		} else {
			m[letters] = append(ss, str)
		}
	}

	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func groupAnagramsAdvance(strs []string) [][]string {
	// golang 的字典可以用数组为key
	var m = make(map[[26]int][]string, len(strs))

	for _, s := range strs {
		var t [26]int
		for i := 0; i < len(s); i++ {
			t[s[i]-'a']++
		}

		m[t] = append(m[t], s)
	}

	var res = make([][]string, len(m))
	var i int
	for _, ss := range m {
		res[i] = ss
		i++
	}

	return res
}

func groupAnagramsSort(strs []string) [][]string {
	mp := map[string][]string{}
	for _, str := range strs {
		s := toBytes(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := toString(s)
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

func toString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func toBytes(s string) []byte {
	var x = (*[2]uintptr)(unsafe.Pointer(&s))
	return *(*[]byte)(unsafe.Pointer(&[3]uintptr{x[0], x[1], x[1]}))
}

func main() {
}
