package main

import "fmt"

func getHint(secret string, guess string) string {
	var set [10]int

	var ac, bc int
	for i := range secret {
		if secret[i] == guess[i] {
			ac++
		} else {
			set[secret[i]-'0']++
		}
	}
	for i := range guess {
		if guess[i] == secret[i] {
			continue
		}
		if set[guess[i]-'0'] > 0 {
			bc++
			set[guess[i]-'0']--
		}
	}

	return fmt.Sprintf("%dA%dB", ac, bc)
}

func getHint2(secret string, guess string) string {
	var set [10]int

	var ac, bc int
	for i := range secret {
		if secret[i] == guess[i] {
			ac++
		} else {
			// 如果当前集合中存在目标字符, 那么就增加一个移位计数
			if set[guess[i]-'0'] > 0 {
				bc++
			}
			// 消耗一个目标字符. 不管是否存在, 直接减去即可
			set[guess[i]-'0']--

			// 如果当前集合缺少目标字符, 那么就增加一个移位计数
			if set[secret[i]-'0'] < 0 {
				bc++
			}
			// 不管是否存在, 都要直接加上去
			set[secret[i]-'0']++
		}
	}

	return fmt.Sprintf("%dA%dB", ac, bc)
}
