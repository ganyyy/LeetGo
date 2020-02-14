package main

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

func main() {
}
