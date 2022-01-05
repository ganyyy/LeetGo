package main

import "unsafe"

func modifyString(s string) string {
	var ret = []byte(s)

	for i, v := range ret {
		if v != '?' {
			continue
		}
		v = 'a'
		if i > 0 {
			v = ((ret[i-1] - 'a' + 1) % 26) + 'a'
		}
		if i < len(ret)-1 && ret[i+1] != '?' {
			if v == ret[i+1] {
				v = ((ret[i+1] - 'a' + 1) % 26) + 'a'
			}
		}
		ret[i] = v
	}

	return *(*string)(unsafe.Pointer(&ret))
}

func main() {
	println(modifyString("??yw?ipkj?"))
}
