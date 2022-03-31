package main

func selfDividingNumbers(left int, right int) []int {
	var check = func(v int) bool {
		var old = v
		for v > 0 {
			var i = v % 10
			if i == 0 || old%i != 0 {
				return false
			}
			v /= 10
		}
		return true
	}

	var ret = make([]int, 0, (right-left)>>4)

	for i := left; i <= right; i++ {
		if !check(i) {
			continue
		}
		ret = append(ret, i)
	}
	return ret
}
