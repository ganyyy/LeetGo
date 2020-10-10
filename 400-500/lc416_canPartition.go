package main

func canPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}
	// 只有是偶数的情况下才可能分为两部分
	if sum&1 != 0 {
		return false
	}
	var mid = sum >> 1
	// 草, 是一个背包问题.
	// 总的容量是 mid, 然后从 nums中选取任意数, 如果存在 dp[mid], 则可以认为是可以分割的
	// 一共存在 len(nums)个物品, 背包的容量是 mid
	var dp = make([]bool, mid+1)
	for i := 0; i < len(nums); i++ {
		for s := mid; s >= nums[i]; s-- {
			if i == 0 {
				dp[s] = s == nums[i]
			} else {
				dp[s] = dp[s] || dp[s-nums[i]]
			}
		}
	}

	return dp[mid]
}

func main() {

}
