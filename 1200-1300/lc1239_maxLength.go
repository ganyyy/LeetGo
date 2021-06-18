package main

func maxLength(arr []string) int {
	// var nums = make([]int, len(arr))

	// for i, s := range arr {
	//     var v int
	//     for j := range s {
	//         v |= int(s[j]-'a')<<1
	//     }
	//     nums[i] = v
	// }

	// 将nums中的数字计算最大的位数和,
	// 反正最大也就26

	// DP? 二叉树?

	// 需要有一种组合的方式

	// 向量图?

	// 每一位可选的对象
	// var valid = make([]int, len(arr))

	// for i := range nums {
	//     var v int
	//     for j := range nums {
	//         if i == j {
	//             continue
	//         }
	//         if nums[i] & nums[j] == 0 {
	//             v |= j<<1
	//         }
	//     }
	//     valid[i] = v
	// }

	// 先确定每一个数字可以连接的对象

	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var back func(i, m int) int

	back = func(i, m int) int {
		if i == len(arr) {
			return 0
		}

		var t = m

		for _, c := range arr[i] {
			if m&(1<<(c-'a')) != 0 {
				// 如果存在任意一位重复的, 直接跳过
				return back(i+1, t)
			}

			m |= 1 << (c - 'a')
		}
		// 选择当前值 or 不选择当前值中, 取较大的那一个
		return max(len(arr[i])+back(i+1, m), back(i+1, t))
	}

	return back(0, 0)
}
