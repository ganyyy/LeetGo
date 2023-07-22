package main

func lemonadeChange(bills []int) bool {
	if len(bills) == 0 {
		return true
	}

	// 第一个必须是5, 否则无法进行下一步
	if bills[0] != 5 {
		return false
	}

	// 如果是10的话, 至少要有一个5
	// 如果是20的话, 至少要有3个5或者1个10+1个5

	var a, b int
	for _, v := range bills {
		// 10的话, 至少要存在一个5
		if v == 10 {
			if a < 1 {
				return false
			}
			a--
			b++
		} else if v == 20 {
			// 优先消耗10, 保留5
			if a >= 1 && b >= 1 {
				a--
				b--
			} else if a >= 3 {
				a -= 3
			} else {
				// 不满足条件, 返回false
				return false
			}
		} else {
			a++
		}
	}
	// 能正确遍历到结尾, 返回false
	return true
}

func lemonadeChange2(bills []int) bool {
	var c5, c10, c20 int
	if len(bills) == 0 || bills[0] != 5 {
		return false
	}
	c5++
	for _, bill := range bills[1:] {
		switch bill {
		case 5:
			c5++
		case 10:
			if c5 < 1 {
				return false
			}
			c10++
			c5--
		case 20:
			need := 15
			if c10 > 0 {
				c10--
				need -= 10
			}
			cnt := need / 5
			if c5 < cnt {
				return false
			}
			c5 -= cnt
			c20++
		default:
			return false
		}
	}
	return true
}
