package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
	if len(name) == len(typed) {
		return name == typed
	}
	if len(typed) < len(name) {
		return false
	}
	if len(name) == 0 {
		return false
	}

	var left, right int

	for left < len(name) && right < len(typed) {
		if name[left] == typed[right] {
			left++
			right++
		} else if right > 0 && typed[right] == typed[right-1] {
			right++
		} else {
			return false
		}
	}

	// typed 剩余的部分
	for right > 0 && right < len(typed) {
		if typed[right] != typed[right-1] {
			return false
		}
		right++
	}
	return left == len(name)
}

func main() {
	var testCases = [][2]string{
		{"alex", "aaleex"},
		{"saeed", "ssaaedd"},
		{"leelee", "lleeelee"},
		{"laiden", "laiden"},
		{"aaa", "aaaa"},
		{"pyplrz", "ppyypllr"},
	}

	for _, c := range testCases {
		fmt.Println(isLongPressedName(c[0], c[1]))
	}
}
