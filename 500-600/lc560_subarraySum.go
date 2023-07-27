package main

func subarraySum(nums []int, k int) int {
	m := make(map[int]int)
	// 表示正好为0的情况
	m[0] = 1
	// 当前的累加和, 返回的结果值
	var sum, ret int
	// 遍历数组的每一个值
	for _, v := range nums {
		// 累加和一直加即可
		sum += v
		// 可以这么理解
		//  如果存在[0:i] 的和(sum-k), 那么 [i+1:cur] 的和就是k
		ret += m[sum-k]
		// 当前位置的个数+1
		m[sum] += 1
	}
	// Mark
	return ret
}

func main() {

}
