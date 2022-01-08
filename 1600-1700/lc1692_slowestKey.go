package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var cnt [26]int

	var parseIndex = func(b byte) int {
		return int(b - 'a')
	}

	cnt[parseIndex(keysPressed[0])] = releaseTimes[0]

	for i := 1; i < len(releaseTimes); i++ {
		var idx = parseIndex(keysPressed[i])
		var pressTime = releaseTimes[i] - releaseTimes[i-1]
		if pressTime > cnt[idx] {
			cnt[idx] = pressTime
		}
	}

	var ret byte
	var max int
	for i := 0; i < len(cnt); i++ {
		if cnt[i] >= max {
			ret = byte(i + 'a')
			max = cnt[i]
		}
	}

	return ret
}
