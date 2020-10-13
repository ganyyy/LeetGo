package main

func commonChars(A []string) []string {
	if len(A) == 0 {
		return nil
	}
	// 求交集
	var res [26]int
	for _, c := range A[0] {
		res[int(c)-'a']++
	}
	// 遍历剩下的
	for _, s := range A[1:] {
		// 统计
		var tmp [26]int
		for _, c := range s {
			tmp[int(c)-'a']++
		}
		// 取最小值
		for i := 0; i < 26; i++ {
			res[i] = min(res[i], tmp[i])
		}
	}
	// 整合结果
	var final []string
	for c, cnt := range res {
		c += 'a'
		for i := 0; i < cnt; i++ {
			final = append(final, string(rune(c)))
		}
	}
	return final
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
