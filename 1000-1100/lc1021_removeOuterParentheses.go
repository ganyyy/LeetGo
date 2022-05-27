package main

func removeOuterParentheses(s string) string {
	var split []int
	var ret = make([]byte, 0, len(s))

	var cnt int
	var start int
	for i := range s {
		if s[i] == '(' {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			split = append(split, i)
			ret = append(ret, s[start+1:i]...)
			start = i + 1
		}
	}
	// fmt.Println(split)
	// var ret = make([]byte, 0, len(s)-2*len(split))
	// for i := 1; i < len(s) && len(split) > 0; i++ {
	//     if i == split[0] {
	//         split = split[1:]
	//         i++
	//         continue
	//     }
	//     ret = append(ret, s[i])
	// }
	return string(ret)
}
