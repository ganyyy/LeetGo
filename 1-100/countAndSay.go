package main

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	pre := countAndSay(n - 1)

	head := pre[0]
	count := 1
	res := make([]byte, 0)
	for i := 1; i < len(pre); i++ {
		if pre[i] != head {
			// X个X
			res = append(res, byte(count)+'0', head)
			head = pre[i]
			count = 1
		} else {
			count++
		}
	}
	res = append(res, byte(count)+'0', head)
	return string(res)
}

func main() {
	fmt.Println(countAndSay(5))
}
