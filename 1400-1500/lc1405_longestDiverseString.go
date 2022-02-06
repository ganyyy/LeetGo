package main

func longestDiverseString(a int, b int, c int) string {

	var tmp = [3]int{a, b, c}

	var max = func(except byte) byte {
		var idx = -1
		var ch byte
		for i := range tmp {
			var b2 = 'a' + byte(i)
			var cnt = tmp[i]
			if cnt == 0 || b2 == except {
				continue
			}
			if idx == -1 || cnt > tmp[idx] {
				idx = i
				ch = b2
			}
		}
		if idx == -1 {
			return 0
		}
		tmp[idx] -= 1
		return ch
	}

	var ret = make([]byte, 0, a+b+c)

	var preA, preB byte
	for {
		var pre byte
		if preA != 0 && preA == preB {
			pre = preA
		}
		var b = max(pre)
		if b == 0 {
			break
		}
		preA, preB = preB, b
		ret = append(ret, b)
	}

	return string(ret)
}

func main() {
	println(longestDiverseString(0, 8, 11))
}
