package main

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

func main() {

}
