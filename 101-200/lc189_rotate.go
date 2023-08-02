//go:build ignore

package main

func rotate2(nums []int, k int) {
	// 循环解法... 这个出问题了
	var ln = len(nums)
	if ln == 0 || k == 0 || k == ln {
		return
	}
	var pre, cur = nums[0], 0
	if ln%k != 0 {
		// 取余结果不为0的情况下, 只需要遍历一轮就行
		for i := 0; i < ln; i++ {
			var next = (cur + k) % ln
			nums[next], pre = pre, nums[next]
			cur = next
		}
	} else {
		// 否则就遍历n轮
		// 比如 ln = 6, k = 4时就无法解决这个问题
		for i := 0; i < k; i++ {
			pre, cur = nums[i], i
			for j := 0; j < ln/k; j++ {
				var next = (cur + k) % ln
				nums[next], pre = pre, nums[next]
				cur = next
			}
		}
	}
}

// 反转解法
func rotate(nums []int, k int) {
	// 防止数组越界
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}
