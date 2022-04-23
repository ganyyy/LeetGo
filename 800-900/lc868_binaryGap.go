package main

func binaryGap(n int) (ret int) {
	var cnt int
	var found bool
	for n != 0 {
		if n&1 == 1 {
			if found {
				if cnt > ret {
					ret = cnt
				}
			}
			cnt = 1
			found = true
		} else {
			cnt++
		}
		n >>= 1
	}

	return ret
}
