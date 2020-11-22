package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var sa [26]int

	for i := 0; i < len(s); i++ {
		sa[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		if l := sa[t[i]-'a']; l > 0 {
			sa[t[i]-'a'] = l - 1
		} else {
			return false
		}
	}

	return true

}
