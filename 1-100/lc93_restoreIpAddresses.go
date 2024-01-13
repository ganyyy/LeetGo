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

func restoreIpAddresses3(s string) []string {
	length := len(s)
	if length > 12 {
		return nil
	}

	// 三个点
	var buf = make([]byte, 0, length+3)
	var ret []string
	var validIP func(buf []byte, s string, remain int)
	validIP = func(buf []byte, s string, remain int) {
		length := len(s)
		if length > remain*3 || length < remain {
			// 超过了, 或者不够
			return
		}
		if remain == 0 && length == 0 {
			ret = append(ret, string(buf))
			return
		}
		if len(buf) > 0 {
			buf = append(buf, '.')
		}
		var val int
		for i := 1; i <= min(length, 3); i++ {
			val = val*10 + int(s[i-1]-'0')
			if i > 1 && val < 10 {
				// 存在前缀0的情况
				break
			}
			if val > 255 {
				// 三位数, 但是和超过了255
				break
			}
			validIP(append(buf, s[:i]...), s[i:], remain-1)
		}
	}
	validIP(buf, s, 4)
	return ret
}

func main() {
	restoreIpAddresses("25525511135")
}
