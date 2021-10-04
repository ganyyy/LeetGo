package main

import (
	"bytes"
	"unsafe"
)

func licenseKeyFormattingBad(s string, k int) string {

	var cnt int
	for i := range s {
		if s[i] != '-' {
			cnt++
		}
	}

	var ret = make([]byte, 0, cnt+(cnt+k-1)/k)

	var idx int
	var num = cnt % k
end:
	for {
		for i := 0; i < num; i++ {
			for idx < len(s) {
				if s[idx] == '-' {
					idx++
					continue
				}
				ret = append(ret, upper(s[idx]))
				idx++
				break
			}
			if idx == len(s) {
				break end
			}
		}
		if num != 0 {
			ret = append(ret, '-')
		}
		num = k
	}
	if len(ret) == 0 {
		return ""
	}
	if ret[len(ret)-1] == '-' {
		ret = ret[:len(ret)-1]
	}
	return *(*string)(unsafe.Pointer(&ret))
}
func upper(s byte) byte {
	if s >= 'a' && s <= 'z' {
		return s + 'A' - 'a'
	}
	return s
}

func licenseKeyFormatting(s string, k int) string {
	// 倒着向前插
	if k <= 0 {
		return ""
	}

	var sb bytes.Buffer
	var cnt int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		sb.WriteByte(upper(s[i]))
		cnt++
		if cnt < k {
			continue
		}
		sb.WriteByte('-')
		cnt = 0
	}
	// 逆序
	var bs = sb.Bytes()
	if len(bs) <= 0 {
		return ""
	}
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	if bs[0] == '-' {
		bs = bs[1:]
	}
	return *(*string)(unsafe.Pointer(&bs))
}

func main() {
	println(licenseKeyFormatting("-a-a-", 1))
}
