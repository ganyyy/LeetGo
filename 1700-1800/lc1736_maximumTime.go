package main

func maximumTime(time string) string {
	var h1 = time[0]
	var h2 = time[1]
	var m1 = time[3]
	var m2 = time[4]
	var ret = []byte(time)

	if h1 == '?' && h2 == '?' {
		ret[0], ret[1] = '2', '3'
	} else if h1 == '?' {
		if h2 < '4' {
			ret[0] = '2'
		} else {
			ret[0] = '1'
		}
	} else if h2 == '?' {
		if h1 == '2' {
			ret[1] = '3'
		} else {
			ret[1] = '9'
		}
	}

	if m2 == '?' {
		ret[4] = '9'
	}

	if m1 == '?' {
		ret[3] = '5'
	}

	return string(ret[:])
}
