package main

func validPalindrome(s string) bool {
	var d bool
	var check func(s string) bool

	check = func(s string) bool {
		for left, right := 0, len(s)-1; left < right; {
			if s[left] != s[right] {
				if d {
					return false
				} else {
					d = true
					return check(s[left:right]) || check(s[left+1:right+1])
				}
			} else {
				left++
				right--
			}
		}
		return true
	}
	return check(s)
}

func main() {

}
