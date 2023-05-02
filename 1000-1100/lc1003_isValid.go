package main

func isValid(s string) bool {
	var st []rune
	// 有序的栈都可以这么搞
	for _, c := range s {
		// a -> 直接入栈
		// b -> 判断栈顶是不是 a, 然后再入栈
		// c -> 判断栈顶是不是 b
		if c > 'a' {
			// b, c都会触发出栈,
			if len(st) == 0 {
				return false
			}
			top := st[len(st)-1]
			st = st[:len(st)-1]
			if c-top != 1 {
				return false
			}
		}
		// a, b都会触发入栈
		if c < 'c' {
			st = append(st, c)
		}
	}
	return len(st) == 0
}
