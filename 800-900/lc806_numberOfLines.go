package main

func numberOfLines(widths []int, s string) []int {
	var row, cur int

	for i := range s {
		var width = widths[s[i]-'a']
		cur += width
		if cur > 100 {
			row++
			cur = width
		}
	}
	return []int{row + 1, cur}
}
