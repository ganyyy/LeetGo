package main

import "fmt"

var ByteMap [3][52]bool

func init() {
	initByteMap(&ByteMap[0], "qwertyuiop")
	initByteMap(&ByteMap[1], "asdfghjkl")
	initByteMap(&ByteMap[2], "zxcvbnm")
}

func getPos(b byte) int {
	if b >= 'a' && b <= 'z' {
		return int(b) - 'a' + 26
	} else {
		return int(b) - 'A'
	}
}

func initByteMap(bits *[52]bool, s string) {
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			bits[s[i]-'a'+26] = true
			bits[s[i]-'a'] = true
		} else {
			bits[s[i]-'A'+26] = true
			bits[s[i]-'A'] = true
		}
	}
}

func getByteMap(b byte) *[52]bool {
	for i := range ByteMap {
		if ByteMap[i][getPos(b)] {
			return &ByteMap[i]
		}
	}
	return nil
}

func findWords(words []string) []string {
	var retIdx = -1

next:
	for i, s := range words {
		var bits = getByteMap(s[0])
		if bits == nil {
			continue
		}
		for _, b := range s[1:] {
			if !bits[getPos(byte(b))] {
				continue next
			}
		}
		retIdx++
		words[retIdx], words[i] = words[i], words[retIdx]
	}

	return words[:retIdx+1]
}

func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}
