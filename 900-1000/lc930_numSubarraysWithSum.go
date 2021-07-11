package main

func numSubarraysWithSum(nums []int, goal int) int {
	// 无法处理前缀0的情况
	var cur int

	var l, r int
	var ret int
	for ; r < len(nums); r++ {
		cur += nums[r]

		// 不够的情况
		if cur < goal {
			continue
		}
		ret++

		for cur > goal {
			cur -= nums[l]
			ret++
			l++
		}
	}

	return ret
}

func numSubarraysWithSumYes(nums []int, goal int) int {
	// 区域前缀和!!!
	var m = make(map[int]int)
	// 需要注意一下, 初始条件
	// 0表示默认不用加
	m[0] = 1
	var cur int
	var ret int
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		// 如果存在 cur-goal, 说明在区间 [X:i]之间存在一个满足其和为goal的连续子数组
		ret += m[cur-goal]
		m[cur]++
	}

	return ret
}

func numSubarraysWithSum3(nums []int, goal int) int {
	var (
		nums0 []int
		count int
	)
	// nums0[i]表示的是第i个1前的连续0的数量
	for _, num := range nums {
		if num == 1 {
			nums0 = append(nums0, count)
			count = 0
		} else {
			count++
		}
	}
	// 末尾补了一下后缀0的个数, 所以 len(nums0) == count(1)+1
	nums0 = append(nums0, count)
	var ans = 0
	// [i:i+goal] 就是统计的一个区间, 这个区间内的1的个数就是goal
	for i := 0; i+goal < len(nums0); i++ {
		if goal == 0 {
			// 0的话, 相当于区间内的0自由组合
			ans = ans + (nums0[i]+1)*nums0[i]/2
		} else {
			// i+goal 说明从[1:goal]之间的子数组的和为 goal
			// 该区间的前缀0的个数*后缀0的个数
			ans = ans + (nums0[i]+1)*(nums0[i+goal]+1)
		}
	}
	return ans
}

func main() {
	println(numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
}
