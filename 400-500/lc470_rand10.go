package main

func rand7() int {
	return 0
}

func rand10() int {
	// 简单而言, 随机概率

	// 构造出一片随机集合, 保留其中等概率出现的一部分
	// 然后对其进行取余

	// 这里获取的是 [1,7] + [1,7] 的所有的结果, 拒绝 一部分数据保证剩余数字满足等概率出现
	// 这个没啥好说的... 不会就是不会
	var a, b = rand7(), rand7()
	if a > 4 && b < 4 {
		return rand10()
	}
	return (a+b)%10 + 1
}
