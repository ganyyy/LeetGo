package main

func subarraysDivByK(A []int, K int) int {
	// 求前缀和
	// 统计和K取余相同的数字的个数
	// 如果 (s[j+1]-s[i]) %k == 0, A[i:j+1] 就是一个数组合适的数组和
	// 从中找出所有和 K 取余相等的前缀和统计其个数, 然后两两组合 就是 最终的结果

	m := make(map[int]int)
	m[0]++
	// 记录前缀和 和 结果
	var pre, res int
	for _, v := range A {
		pre = (v + pre) % K
		// 这里要防止出现负值
		if pre < 0 {
			pre += K
		}
		// 先加结果, 等同于 C(x, 2)
		res += m[pre]
		m[pre]++
	}
	return res
}

func main() {

}
