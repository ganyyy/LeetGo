package main

import (
	"strconv"
	"strings"
)

func optimalDivision(nums []int) string {
	var sb strings.Builder

	sb.WriteString(strconv.Itoa(nums[0]))

	if len(nums) == 1 {
		return sb.String()
	} else if len(nums) == 2 {
		sb.WriteString("/")
		sb.WriteString(strconv.Itoa(nums[1]))
		return sb.String()
	}

	sb.WriteString("/(")
	sb.WriteString(strconv.Itoa(nums[1]))
	for _, v := range nums[2:] {
		sb.WriteString("/")
		sb.WriteString(strconv.Itoa(v))
	}
	sb.WriteString(")")
	return sb.String()
}
