package main

import "strings"

func generateTheString(n int) string {

	if n&1 == 0 {
		return strings.Repeat("a", n-1) + "b"
	} else {
		return strings.Repeat("a", n)
	}
}
