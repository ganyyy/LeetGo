package main

import (
	"fmt"
	"testing"
)

func Test_LC763PartitionLabels(t *testing.T) {
	var ret = partitionLabels("aaaab")
	t.Log(ret)
}

func TestMonotoneIncreasingDigits(t *testing.T) {
	t.Log(monotoneIncreasingDigits(10))
}

func TestAccountMerge(t *testing.T) {
	var src = [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
	}

	fmt.Println(accountsMerge(src))
}
