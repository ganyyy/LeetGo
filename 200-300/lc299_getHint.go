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
			if set[guess[i]-'0'] > 0 {
				bc++
			}
			set[guess[i]-'0']--
			if set[secret[i]-'0'] < 0 {
				bc++
			}
			set[secret[i]-'0']++
		}
	}

	return fmt.Sprintf("%dA%dB", ac, bc)
}
