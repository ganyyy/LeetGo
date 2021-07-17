package main

var prim = []uint64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}

func groupAnagrams(strs []string) [][]string {
	// i need a hash alg

	// 有一种做法, 是将每个字符串进行排序, 然后如果相等的话, 就塞到一起

	var getSorted = func(s string) uint64 {
		var sum uint64 = 1

		for i := range s {
			sum *= prim[s[i]-'a']
		}
		return sum
	}

	var m = make(map[uint64][]string)

	for _, s := range strs {
		var vs = getSorted(s)
		m[vs] = append(m[vs], s)
	}

	var ret = make([][]string, 0, len(m))

	for _, ss := range m {
		ret = append(ret, ss)
	}

	return ret
}
