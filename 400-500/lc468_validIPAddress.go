package main

import (
	"strconv"
	"strings"
)

func validIPAddress(queryIP string) string {
	const (
		V4   = "IPv4"
		V6   = "IPv6"
		None = "Neither"
	)

	if ss := strings.Split(queryIP, "."); len(ss) == 4 {
		for _, s := range ss {
			if len(s) > 1 && s[0] == '0' {
				return None
			}
			if v, err := strconv.Atoi(s); err != nil || v > 255 {
				return None
			}
		}
		return V4
	} else if ss = strings.Split(queryIP, ":"); len(ss) == 8 {
		for _, s := range ss {
			if len(s) > 4 {
				return None
			}
			if _, err := strconv.ParseUint(s, 16, 64); err != nil {
				return None
			}
		}
		return V6
	}
	return None
}
