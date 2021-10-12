package main

import "strconv"

func fizzBuzz(n int) []string {
	var ret = make([]string, 0, n)
	var n3, n5 = 3, 5
	const s3, s5 = "Fizz", "Buzz"
	for i := 1; i <= n; i++ {
		var tmp string
		if i == n3 {
			n3 += 3
			tmp += s3
		}
		if i == n5 {
			n5 += 5
			tmp += s5
		}
		if tmp == "" {
			tmp = strconv.Itoa(i)
		}
		ret = append(ret, tmp)
	}

	return ret
}
