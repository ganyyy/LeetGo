package main

func checkSubarraySum(nums []int, k int) bool {
	// // 前缀和

	// 超时1 93/94
	// // 非负数组, 和肯定是递增的.

	// // 判断是否满足条件, 看的是

	// var tmp = make([]int, len(nums)+1)
	// for i := 0; i < len(nums); i++ {
	//     tmp[i+1] += nums[i]+tmp[i]
	// }

	// // O(n^2)级别的时间复杂度, 或者可以搞个双指针啥的.. 貌似也不用遍历到末尾欸
	// // fmt.Println(tmp)
	// for i := 1; i < len(tmp)-1; i++ {
	//     for j := i+1; j < len(tmp); j++ {
	//         if (tmp[j]-tmp[i-1]) % k == 0 {
	//             return true
	//         }
	//     }
	// }

	// 错误2 90/94
	// // 统计中间存在空余的数字
	// var m = make(map[int]bool)
	// // 统计本身和k取余为0的连续数字的个数
	// var preCnt int
	// for i := 1; i < len(nums); i++ {
	//     if nums[i] % k == 0 {
	//         // 连续直接取余为零的情况
	//         preCnt++
	//         if preCnt >= 2 {
	//             return true
	//         }
	//     } else {
	//         // 重置一下, 中间断了
	//         preCnt = 0
	//     }
	//     nums[i] += nums[i-1]
	//     nums[i] %= k
	//     // 余数为0, 或者前边存在+当前数为k的数字
	//     if nums[i] == 0 || m[k-nums[i]] {
	//         return true
	//     }
	//     // 记录当前保留的数字
	//     m[nums[i]] = true
	// }

	// 提前终止条件
	if k == 0 || len(nums) < 2 {
		return false
	}
	if k < 0 {
		k = -k
	}

	// 前缀和 %k 值和其对应的索引
	var m = make(map[int]int)
	// 增加一个默认值
	m[0] = -1

	// 同余关系:
	// sum[:x] % k == sum[:y] % k
	// 说明 [x:y] 是一个满足条件的子数组.

	var sum int
	for i, v := range nums {
		if i > 0 && v == 0 && nums[i-1] == 0 {
			return true
		}
		sum += v
		t := sum % k
		if idx, ok := m[t]; ok {
			// 存在并且相差超过1, 返回true
			if i-idx > 1 {
				return true
			}
		} else {
			m[t] = i
		}
	}

	return false
}
