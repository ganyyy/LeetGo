//go:build ignore

package main

import (
	"strconv"
)

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return true
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return nil
}

func parse(s string) int {
	var ret, _ = strconv.Atoi(s)
	return ret
}

func deserialize(s string) (nested *NestedInteger) {
	nested = &NestedInteger{}
	if len(s) == 0 {
		return
	}
	if s[0] != '[' {
		nested.SetInteger(parse(s))
		return
	}
	if len(s) <= 2 {
		return
	}

	var start, cnt = 1, 0
	for i := 1; i < len(s); i++ {
		if cnt == 0 && (s[i] == ',' || i == len(s)-1) {
			nested.Add(*deserialize(s[start:i]))
			start = i + 1
		} else if s[i] == '[' {
			cnt++
		} else if s[i] == ']' {
			cnt--
		}
	}
	return
}
func main() {
	deserialize("[123,[456,[789], 888], 999]")
}
