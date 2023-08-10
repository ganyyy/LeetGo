package main

import "strconv"

func restoreIpAddresses(s string) []string {
	if len(s) > 12 {
		return nil
	}

	var check func([]byte, []byte, int)

	var res []string

	check = func(pre, s []byte, remain int) {
		// 超长了, 返回
		if len(s) > remain*3 {
			return
		}
		ln := len(s)
		// 都清零了, 可以加一个结果
		if ln == 0 && remain == 0 {
			res = append(res, string(pre[:len(pre)-1]))
		}
		// 不够了, 返回
		if ln < remain {
			return
		}
		for i := 1; i <= min(ln, 3); i++ {
			// 去掉前边的0
			if v, _ := strconv.Atoi(string(s[:i])); v > 255 {
				continue
			} else if i > 2 && v < 100 {
				continue
			} else if i > 1 && v < 10 {
				continue
			}
			check(append([]byte(nil), append(append(pre, s[:i]...), '.')...), s[i:], remain-1)
		}
	}
	check(nil, []byte(s), 4)
	return res
}

func main() {
	restoreIpAddresses("25525511135")
}
