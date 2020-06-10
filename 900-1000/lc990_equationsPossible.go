package main

import "fmt"

func equationsPossible(equations []string) bool {
	var option [26]byte
	for i := 0; i < 26; i++ {
		option[i] = byte(i)
	}

	var find func(byte) byte

	find = func(a byte) byte {
		for option[a] != a {
			a = option[a]
		}
		return a
	}

	// 记录不相等的
	var lst [][2]byte

	for _, v := range equations {
		a, b := v[0]-'a', v[3]-'a'
		if a != b {
			if v[1] == '=' {
				x, y := find(a), find(b)
				if x != y {
					option[y] = x
				}
			} else {
				lst = append(lst, [2]byte{a, b})
			}
		} else {
			if v[1] == '!' {
				return false
			}
		}
	}

	// 从头到尾在来一遍, 找出不同的
	for _, v := range lst {
		a, b := v[0], v[1]
		if find(a) == find(b) {
			return false
		}
	}

	return true
}

func main() {
	equations := []string{
		"a==b", "b!=a",
	}
	fmt.Println(equationsPossible(equations))
}
