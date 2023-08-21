package main

func canChange(start string, target string) bool {
	i, j, n := 0, 0, len(start)
	for i < n && j < n {
		for i < n && start[i] == '_' {
			i++
		}
		for j < n && target[j] == '_' {
			j++
		}
		if i < n && j < n {
			if start[i] != target[j] {
				return false
			}
			c := start[i]
			// 当字符是L时, start 所处的下标一定要大于target, 因为start中的L是要往左移动的
			// 当字符是R时, start 所处的下标一定要小于target, 因为start中的R是要往右移动的
			if c == 'L' && i < j || c == 'R' && i > j {
				return false
			}
			i++
			j++
		}
	}
	for i < n {
		// target提前结束了
		// 比如 _L__ 和 ___L
		if start[i] != '_' {
			return false
		}
		i++
	}
	for j < n {
		// start提前结束了
		// 比如 ___L 和 _L__
		if target[j] != '_' {
			return false
		}
		j++
	}
	return true
}
