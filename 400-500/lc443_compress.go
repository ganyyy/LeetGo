package main

import (
	"bytes"
	"strconv"
)

func compressUseBuffer(chars []byte) int {
	var tmp = bytes.NewBuffer(chars)

	// 双指针啊

	var l, r int

	var writeToBuffer = func() {
		tmp.WriteByte(chars[l])
		if r-l > 1 {
			tmp.WriteString(strconv.Itoa(r - l))
		}
	}

	for r < len(chars) {
		if chars[r] != chars[l] {
			writeToBuffer()
		}
		r++
	}

	if r > l {
		writeToBuffer()
	}
	copy(chars, tmp.Bytes())
	return tmp.Len()
}

func compress(chars []byte) int {

	// 双指针啊

	var base, l, r int

	var writeToBuffer = func() {
		chars[base] = chars[l]
		if r-l > 1 {
			var s = strconv.Itoa(r - l)
			for i := range s {
				base++
				chars[base] = s[i]
			}
		}
		base++
	}

	for r < len(chars) {
		if chars[r] != chars[l] {
			writeToBuffer()
			l = r
		}
		r++
	}

	if r > l {
		writeToBuffer()
	}
	return base
}
