//go:build ignore
// +build ignore

package main

import "strings"

func replaceSpace(s string) string {
	var ss = strings.Split(s, " ")
	return strings.Join(ss, "%20")
}
