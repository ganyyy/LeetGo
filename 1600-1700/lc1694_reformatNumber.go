package main

import "bytes"

func reformatNumber(number string) string {
	buf := make([]byte, 0, len(number))
	for _, n := range number {
		c := byte(n)
		if c == '-' || c == ' ' {
			continue
		}
		buf = append(buf, c)
	}
	ret := bytes.NewBuffer(nil)

	fill := func(n int) {
		for i := 0; i < len(buf)-n; i += 3 {
			ret.Write(buf[i : i+3])
			ret.WriteByte('-')
		}
	}

	n := len(buf)
	switch n % 3 {
	case 1:
		if n > 3 {
			// 4
			fill(4)
			ret.Write(buf[n-4 : n-2])
			ret.WriteByte('-')
			ret.Write(buf[n-2 : n])
		} else {
			// 1
			fill(1)
			ret.Write(buf[n-1:])
		}
	case 2:
		fill(2)
		ret.Write(buf[n-2:])
	default:
		fill(3)
		ret.Write(buf[n-3:])
	}

	return ret.String()
}
