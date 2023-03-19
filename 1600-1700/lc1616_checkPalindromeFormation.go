package main

func checkPalindromeFormation(a, b string) bool {
	return checkConcatenation(a, b) || checkConcatenation(b, a)
}

func checkConcatenation(a, b string) bool {
	left, right := 0, len(a)-1
	for left < right && a[left] == b[right] {
		left++
		right--
	}
	// 整个就是一个互为回文的例子
	if left >= right {
		return true
	}
	// 比如吧
	// abc|M...|...
	// ...|...N|cba
	// 此时, 只需要 a中的 M... 是一个回文, 或者 b中的 ...N是一个回文 就行了
	return checkSelfPalindrome(a[left:right+1]) || checkSelfPalindrome(b[left:right+1])
}

func checkSelfPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		left++
		right--
	}
	return left >= right
}
