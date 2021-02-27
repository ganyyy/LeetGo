package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longestSubstring(s string, k int) int {

	var n = len(s)

	if n < k {
		return 0
	}

	var cnt [26]int
	for i := range s {
		cnt[s[i]-'a']++
	}

	var l, r, res int

	for ; r < n; r++ {
		// 如果字符 s[r] 出现的次数小于 k, 说明这个字符串一定不会包含在结果种
		// 直接进行切割比较即可
		if cnt[s[r]-'a'] < k {
			res = max(res, longestSubstring(s[l:r], k))
			l = r + 1
		}
	}

	// 如果 l == 0, 说明所有的字符都满足大于k, 直接返回即可
	if l == 0 {
		return n
	}
	// 否则就要从 res 和 l:n(r) 中计算一下最大值
	return max(res, longestSubstring(s[l:n], k))

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var sort = func(nums []int) {
		var cnt int
		var m = len(nums)
		for i := 0; i < len(nums); i++ {
			for nums[i] != m-i {
				nums[i], nums[m-nums[i]] = nums[m-nums[i]], nums[i]
			}
		}
		fmt.Println(cnt)
	}

	var num = make([]int, 10)
	for i := 0; i < 10; i++ {
		num[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(10, func(i, j int) {
		num[i], num[j] = num[j], num[i]
	})
	fmt.Println(num)
	sort(num)
	fmt.Println(num)
}
