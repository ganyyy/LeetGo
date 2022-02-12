package main

import "math"

func maxNumberOfBalloons(text string) int {
	var cnt [5]int
	for _, b := range text {
		switch b {
		case 'b':
			cnt[0]++
		case 'a':
			cnt[1]++
		case 'l':
			cnt[2]++
		case 'o':
			cnt[3]++
		case 'n':
			cnt[4]++
		}
	}

	var ret = math.MaxInt32
	cnt[2] /= 2
	cnt[3] /= 2
	for _, v := range cnt {
		if ret > v {
			ret = v
		}
	}

	return ret
}
