package main

func asteroidCollision(asteroids []int) []int {
	var stack []int
next:
	for _, v := range asteroids {
		if v >= 0 {
			stack = append(stack, v)
			continue
		}
		for len(stack) > 0 {
			var last = stack[len(stack)-1]
			if last > 0 {
				if abs(v) > last {
					// 不停的弹出顶端
					stack = stack[:len(stack)-1]
					continue
				} else if abs(v) == last {
					// 碰没了
					stack = stack[:len(stack)-1]
					continue next
				}
				// last > abs(v), 直接看下一个
				continue next
			} else {
				// 都是小于0的, 添加到栈中
				stack = append(stack, v)
				continue next
			}
		}
		// 有一种可能是
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
	}
	return stack
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
