package main

func isOneBitCharacterBad(bits []int) bool {
	// 去掉最后一位, 看看能不能拼接成功

	// 想想不符合规范的是什么样的

	// 从后向前匹配

	// 如果是0, 那么存在两种情况

	// 如果是1, 那么只有一种情况

	return check(bits[:len(bits)-1])
}

func check(bits []int) bool {
	if len(bits) == 0 {
		return true
	}
	if len(bits) == 1 {
		return bits[0] == 0
	}
	var pre, last = bits[len(bits)-2], bits[len(bits)-1]

	if pre == 0 {
		// 只能是单个字符
		return last == 0 && check(bits[:len(bits)-1])
	} else {
		return check(bits[:len(bits)-2]) || (last == 0 && check(bits[:len(bits)-1]))
	}
}

func isOneBitCharacter(bits []int) bool {
	// 为0的情况下, 吃1位

	// 为1的情况下, 吃2位

	var start int

	for start < len(bits)-1 {
		start += bits[start] + 1
	}
	return start == len(bits)-1

}
