package main

func removePalindromeSub(s string) int {
	// 如果是回文串, 那就是1次, 否则就是两次(a, b)

	var left, right = 0, len(s) - 1

	for left < right {
		if s[left] != s[right] {
			break
		}
		left++
		right--
	}

	if left >= right {
		return 1
	}
	// 先删除a, 在删除b
	return 2

}
