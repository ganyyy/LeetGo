//go:build ignore

package main

import "math"

func maxValueAfterReverse(nums []int) int {
	base, d, n := 0, 0, len(nums)

	// 假设 1 <= i < j < n-1
	// a = nums[i-1], b = nums[i], x = nums[j], y = nums[j+1]
	// 很明显, 反转一个子数组, 只和他的边界的变化有关系! 因为中间的绝对值和并没有发生任何变化!
	// 翻转前是 |a-b| + |x-y|, 反转后是 |a-x| + |b-y|
	// 所以差值就是
	// d=∣a−x∣+∣b−y∣−∣a−b∣−∣x−y∣ ①

	// 三个式子:
	// a+b=max(a,b)+min(a,b)
	// a+b+∣a−b∣=2⋅max(a,b)
	// a+b−∣a−b∣=2⋅min(a,b)

	// 三种情况:
	// 1. max(a,b)≤min(x,y) | min(x,y)≤max(a,b)  带入到①中 -> d = 2⋅min(x,y)−2⋅max(a,b) >= 0
	// 2. max(a,x)≤min(b,y) | min(b,y)≤max(a,x)  带入到①中 -> d <= 0
	// 3. max(a,y)≤min(b,x) | min(a,y)≤max(b,x)  带入到①中 -> d == 0

	// mx: min(a,b)的最大值
	// mn: max(a,b)的最小值
	mx, mn := math.MinInt, math.MaxInt

	for i := 1; i < n; i++ {
		// 这里的a, b也可以看成是 x,y
		a, b := nums[i-1], nums[i]
		base += abs(a - b)
		mx = max(mx, min(a, b))
		mn = min(mn, max(a, b))
		// 这里的作用是?
		// 如果反转的左边界是左端点, 则可以省去a(-1), 此时就变成了 |b(a+1=0)-x| - |x-y|
		// 如果反转的右边界是右端点, 则可以省去y(n),  此时就变成了 |a-x(y-1=n-1)| - |a-b|
		d = max(d, max(abs(nums[0]-b)-abs(a-b), // i=0
			abs(nums[n-1]-a)-abs(a-b))) // j=n-1
	}
	return base + max(d, 2*(mx-mn))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
