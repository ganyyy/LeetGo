package main

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for !checkValid(s[left]) && left < right {
			left++
		}
		for !checkValid(s[right]) && left < right {
			right--
		}
		if getVal(s[left]) != getVal(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func getVal(a byte) byte {
	if a >= 'a' {
		a -= 'a'
	} else if a >= 'A' {
		a -= 'A'
	}
	return a
}

func checkValid(b byte) bool {
	return (b >= '0' && b <= '9') ||
		(b >= 'a' && b <= 'z') ||
		(b >= 'A' && b <= 'Z')
}
