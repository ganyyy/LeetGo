package main

func minMutation(start string, end string, bank []string) int {

	var queue []string

	var checkNext = func(src string) (ret []string) {
		for i, b := range bank {
			if b == "" {
				continue
			}
			// 差异值
			var diff int
			for idx := range b {
				if src[idx] == b[idx] {
					continue
				}
				diff++
			}

			if diff == 1 {
				// 差异值为1, 可以进行转换, 同时将bank中的值置空标记为已经使用过
				bank[i] = ""
				ret = append(ret, b)
			}
		}
		return
	}
	var round int
	queue = checkNext(start)
	for len(queue) != 0 {
		round++
		var idx = len(queue)
		for i := 0; i < idx; i++ {
			if queue[i] == end {
				return round
			}
			queue = append(queue, checkNext(queue[i])...)
		}
		queue = queue[idx:]
	}
	return -1
}
