//go:build ignore

package main

func maxSumTwoNoOverlap(nums []int, firstLen, secondLen int) (ans int) {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x // 计算 nums 的前缀和
	}
	f := func(firstLen, secondLen int) {
		maxSumA := 0
		for i := firstLen + secondLen; i <= n; i++ {
			// first 的前缀和 s[i-secondLen]-s[i-secondLen-firstLen]
			maxSumA = max(maxSumA, s[i-secondLen]-s[i-secondLen-firstLen])
			// second 的前缀和 s[i]-s[i-secondLen]
			ans = max(ans, maxSumA+s[i]-s[i-secondLen])
		}
	}
	// a可以出现在左边, 也可以出现在右边. 两种情况都需要考虑一下
	f(firstLen, secondLen) // 左 a 右 b
	f(secondLen, firstLen) // 左 b 右 a
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/*
func maxSumTwoNoOverlap(nums []int, firstLen, secondLen int) (ans int) {
    n := len(nums)
    s := make([]int, n+1)
    for i, x := range nums {
        s[i+1] = s[i] + x // 计算 nums 的前缀和
    }
    maxSumA, maxSumB := 0, 0
    for i := firstLen + secondLen; i <= n; i++ {
        maxSumA = max(maxSumA, s[i-secondLen]-s[i-secondLen-firstLen])
        maxSumB = max(maxSumB, s[i-firstLen]-s[i-firstLen-secondLen])
        ans = max(ans, max(maxSumA+s[i]-s[i-secondLen], // 左 a 右 b
                           maxSumB+s[i]-s[i-firstLen])) // 左 b 右 a
    }
    return
}

func max(a, b int) int { if b > a { return b }; return a }
*/
