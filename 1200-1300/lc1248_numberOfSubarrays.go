package main

func numberOfSubarrays(nums []int, k int) int {
	ln := len(nums)
	var res, oddCount int
	// 为啥ln+2呢?因为需要保存开头和结尾
	arr := make([]int, ln+2)
	for i := 0; i < ln; i++ {
		// 给每一个奇数找到他的位置
		if (nums[i] & 1) == 1 {
			oddCount++
			arr[oddCount] = i
		}
	}
	// 左边界
	arr[0] = -1
	// 右边界
	arr[oddCount+1] = ln

	// arr[i]是窗口左边界
	// arr[i+k-1] 是窗口右边界
	// arr[i-1]是左边的上一个奇数，在此之后到arr[i]都可选
	// arr[i+k]是右边的下一个奇数，在此之前都arr[i+k-1]都可选
	//前面可选部分长度为arr[i]-arr[i-1]
	//后面可选部分长度为arr[i+k]-arr[i+k-1]
	//总的可能数等于前后可选的组合

	// 这里的oddCount+2指的是前后头
	for i := 1; i+k < oddCount+2; i++ {
		res += (arr[i] - arr[i-1]) * (arr[i+k] - arr[i+k-1])
	}
	return res
}

func main() {

}
