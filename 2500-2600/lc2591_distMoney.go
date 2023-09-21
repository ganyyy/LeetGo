package main

func distMoney(money int, children int) int {
	if children > money {
		return -1
	}
	// 获取8美元的儿童计数
	mp := money / 8
	for {
		// 剩余的钱的数量
		remainMoney := money - mp*8
		// 剩余人的个数
		remainChildren := children - mp

		if (remainMoney == 4 && remainChildren == 1) ||
			remainMoney < remainChildren ||
			remainChildren < 0 ||
			(remainChildren == 0 && remainMoney > 0) {
			// 特殊情况1: 只剩一个人, 并且还只剩了4刀
			// 特殊情况2: 剩余的钱不够每人一块的
			// 特殊情况3: 人少了
			// 特殊情况4: 钱有剩余的
			mp--
			continue
		}
		break
	}
	return mp
}
