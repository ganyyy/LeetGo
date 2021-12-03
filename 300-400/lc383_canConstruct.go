package main

func canConstruct(ransomNote string, magazine string) bool {
	var a [26]int

	for _, v := range magazine {
		a[v-'a']++
	}

	for _, v := range ransomNote {
		a[v-'a']--
		if a[v-'a'] < 0 {
			return false
		}
	}

	return true
}
