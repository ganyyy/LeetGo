package main

import "unsafe"

func backspaceCompare(S string, T string) bool {
	return toString(filter(S)) == toString(filter(T))
}

func toString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func filter(s string) []byte {
	var t = make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(t) != 0 {
				t = t[:len(t)-1]
			}
		} else {
			t = append(t, s[i])
		}
	}
	return t
}

func main() {

}
