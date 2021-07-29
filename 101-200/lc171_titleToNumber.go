package main

func titleToNumber(columnTitle string) int {
	var ret int

	for i := range columnTitle {
		var c = columnTitle[i]
		ret = ret*26 + int(c-'A') + 1
	}
	return ret
}
