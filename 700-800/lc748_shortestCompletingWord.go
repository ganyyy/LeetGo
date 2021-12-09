package main

func shortestCompletingWord(licensePlate string, words []string) string {
	var licenseSet [26]int

	for i := range licensePlate {
		var c = licensePlate[i]
		switch {
		case c >= 'A' && c <= 'Z':
			licenseSet[c-'A']++
		case c >= 'a' && c <= 'z':
			licenseSet[c-'a']++
		}
	}

	var cnt [26]int
	var minWord = "XXXXXXXXXXXXXXXXX" // 比15长就行

	for _, w := range words {
		for i := range w {
			cnt[w[i]-'a']++
		}
		var find = true
		for i, v := range cnt {
			if v < licenseSet[i] {
				find = false
			}
			cnt[i] = 0
		}
		if find && len(minWord) > len(w) {
			minWord = w
		}
	}

	return minWord
}
