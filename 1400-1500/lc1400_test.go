package main

import (
	"math/rand"
	"testing"
)

var nums = func() []int {
	var res []int
	for i := 0; i < 10000; i++ {
		res = append(res, rand.Intn(19999))
	}
	return res
}()
var limit = 4

func BenchmarkLongestSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestSubarray(nums, limit)
	}
}

func BenchmarkLongestSubArray2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestSubarray2(nums, limit)
	}
}

func BenchmarkLongestSubArray3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		longestSubarray3(nums, limit)
	}
}
