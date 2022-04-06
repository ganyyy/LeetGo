package main

import "strings"

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Index(s+s, goal) != -1
}
