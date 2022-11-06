package main

func getPos(s string) (pos []string) {
	if s[0] != '0' || s == "0" {
		pos = append(pos, s)
	}
	if s[len(s)-1] == '0' {
		// 末尾0不能加小数点
		return
	}
	if s[0] == '0' {
		// 开头0只能在其后面加小数点
		pos = append(pos, s[:1]+"."+s[1:])
		return
	}
	// 其他情况随便加小数点
	for p := 1; p < len(s); p++ {
		pos = append(pos, s[:p]+"."+s[p:])
	}
	return
}

func ambiguousCoordinates(s string) (res []string) {
	n := len(s) - 2
	s = s[1 : len(s)-1]
	for l := 1; l < n; l++ {
		// 左右两边进行分割
		lt := getPos(s[:l])
		if len(lt) == 0 {
			continue
		}
		rt := getPos(s[l:])
		if len(rt) == 0 {
			continue
		}
		for _, i := range lt {
			for _, j := range rt {
				res = append(res, "("+i+", "+j+")")
			}
		}
	}
	return
}
