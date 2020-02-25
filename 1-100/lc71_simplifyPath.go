package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func simplifyPath(path string) string {
	if len(path) == 0 {
		return path
	}
	p := strings.Split(path, "/")
	stack := make([]string, 0)
	for i := 0; i < len(p); i++ {
		switch p[i] {
		case "", ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, p[i])
		}
	}
	build := strings.Builder{}
	if len(stack) != 0 {
		for i := 0; i < len(stack); i++ {
			build.WriteByte('/')
			build.WriteString(stack[i])
		}
	} else {
		build.WriteByte('/')
	}

	return build.String()
}

func simplifyPath2(path string) string {
	return filepath.Clean(path)
}

func main() {
	fmt.Println(simplifyPath2("/a//b////c/d//././/.."))
}
