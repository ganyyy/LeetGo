package main

import "strings"

var num1 = []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
var num2 = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen",
	"Eighteen", "Nineteen"}
var num3 = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	var s string
	var build func(num int) string
	build = func(num int) string {
		if num < 10 {
			s = num1[num]
		} else if num < 20 {
			s = num2[num-10]
		} else if num < 100 {
			s = num3[num/10] + " " + num1[num%10]
		} else if num < 1000 {
			s = num1[num/100] + " Hundred " + build(num%100)
		} else if num < 1000000 {
			s = build(num/1000) + " Thousand " + build(num%1000)
		} else if num < 1000000000 {
			s = build(num/1000000) + " Million " + build(num%1000000)
		} else {
			s = build(num/1000000000) + " Billion " + build(num%1000000000)
		}
		s = strings.TrimSpace(s)

		return s
	}
	return build(num)
}
