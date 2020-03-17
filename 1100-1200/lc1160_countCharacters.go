package main

func countCharacters(words []string, chars string) int {
	lc := len(chars)
	if lc == 0 {
		return 0
	}
	count := [26]byte{}
	for i := 0; i < lc; i++ {
		count[chars[i]-'a']++
	}
	var res int
	for _, word := range words {
		if len(word) > lc {
			continue
		}
		i, lw, bc := 0, len(word), count
		for ; i < lw; i++ {
			if bc[word[i]-'a'] > 0 {
				bc[word[i]-'a']--
			} else {
				break
			}
		}
		if i == lw {
			res += lw
		}
	}
	return res
}

func main() {

}
