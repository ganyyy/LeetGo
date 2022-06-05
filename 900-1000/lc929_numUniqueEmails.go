package main

import "strings"

func numUniqueEmails(emails []string) int {
	var parse = func(str string) string {
		var atIdx = strings.Index(str, "@")
		if atIdx == -1 {
			return str
		}
		var plusIdx = strings.Index(str, "+")
		var replace = str[:atIdx]
		if plusIdx != -1 {
			replace = str[:plusIdx]
		}
		return strings.ReplaceAll(replace, ".", "") + str[atIdx:]
	}
	var count = make(map[string]struct{})
	for _, email := range emails {
		count[parse(email)] = struct{}{}
	}
	return len(count)
}
