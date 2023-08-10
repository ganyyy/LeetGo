package main

import "math"

func rob(nums []int) int {
	// 选第一个和不选第一个, 取最大值?
	ln := len(nums)

	if ln == 0 {
		return 0
	}

	if ln == 1 {
		return nums[0]
	}

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var subRob = func(nums []int) int {
		var a, b int
		for i := 0; i < len(nums); i++ {
			a, b = b, max(nums[i]+a, b)
		}
		return b
	}
	return max(subRob(nums[:ln-1]), subRob(nums[1:]))
}

func rob2(nums []int) int {
	// 主要是开头和结尾的区别
	// 中间的可以认为存在两种处理模式:
	// max(nums[i] + nums[i-2], nums[i-1])

	// 对于小于3的情况, 直接返回最大值
	var res = math.MinInt32
	if len(nums) <= 3 {
		for i := 0; i < len(nums); i++ {
			if nums[i] > res {
				res = nums[i]
			}
		}
		return res
	}

	// 开始处理

	// 想复杂了. 如果从开头就进行大小判断的话, 就没这个问题了
	// 分两趟, 第一趟看[0, n-2], 第二趟看[1, n-1]
	//var a, b, c = nums[0], nums[1], max(nums[1], nums[2]+nums[0])
	//for i := 3; i < len(nums)-1; i++ {
	//	a, b, c = b, c, max(c, max(a+nums[i], b+nums[i]))
	//}
	//res = max(a, max(b, c))
	//a, b, c = nums[1], nums[2], max(nums[2], nums[3]+nums[1])
	//for i := 4; i < len(nums); i++ {
	//	a, b, c = b, c, max(c, max(a+nums[i], b+nums[i]))
	//}
	var a, b int
	for i := 0; i < len(nums)-1; i++ {
		a, b = b, max(a+nums[i], b)
	}
	res = max(a, b)
	a, b = 0, 0
	for i := 1; i < len(nums); i++ {
		a, b = b, max(a+nums[i], b)
	}
	return max(res, max(a, b))
}

func main() {

}
