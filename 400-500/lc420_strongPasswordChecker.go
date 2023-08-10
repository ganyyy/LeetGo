package main

import "unicode"

func strongPasswordChecker(password string) int {
	hasLower, hasUpper, hasDigit := 0, 0, 0
	for _, ch := range password {
		if unicode.IsLower(ch) {
			hasLower = 1
		} else if unicode.IsUpper(ch) {
			hasUpper = 1
		} else if unicode.IsDigit(ch) {
			hasDigit = 1
		}
	}
	// 出现字符的种类
	categories := hasLower + hasUpper + hasDigit

	switch n := len(password); {
	case n < 6:
		// 小于6个, 只需要进行添加
		// 添加的下限就是个数和种类的最大值
		return max(6-n, 3-categories)
	case n <= 20:
		// 6-20, 只需要执行替换逻辑
		replace, cnt, cur := 0, 0, '#'
		for _, ch := range password {
			if ch == cur {
				cnt++
			} else {
				replace += cnt / 3
				cnt = 1
				cur = ch
			}
		}
		replace += cnt / 3
		return max(replace, 3-categories)
	default:
		// 替换次数和删除次数
		replace, remove := 0, n-20
		// k mod 3 = 1 的组数，即删除 2 个字符可以减少 1 次替换操作
		rm2, cnt, cur := 0, 0, '#'
		for _, ch := range password {
			if ch == cur {
				cnt++
				continue
			}
			if remove > 0 && cnt >= 3 {
				if cnt%3 == 0 {
					// 如果是 k % 3 = 0 的组，那么优先删除 1 个字符，减少 1 次替换操作
					remove--
					replace--
				} else if cnt%3 == 1 {
					// 如果是 k % 3 = 1 的组，那么存下来备用
					rm2++
				}
				// k % 3 = 2 的组无需显式考虑
			}
			replace += cnt / 3
			cnt = 1
			cur = ch
		}

		if remove > 0 && cnt >= 3 {
			if cnt%3 == 0 {
				remove--
				replace--
			} else if cnt%3 == 1 {
				rm2++
			}
		}

		replace += cnt / 3

		// 使用 k % 3 = 1 的组的数量，由剩余的替换次数、组数和剩余的删除次数共同决定
		use2 := min(min(replace, rm2), remove/2)
		replace -= use2
		remove -= use2 * 2

		// 由于每有一次替换次数就一定有 3 个连续相同的字符（k / 3 决定），因此这里可以直接计算出使用 k % 3 = 2 的组的数量
		use3 := min(replace, remove/3)
		replace -= use3
		remove -= use3 * 3
		return (n - 20) + max(replace, 3-categories)
	}
}
