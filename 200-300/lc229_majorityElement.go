package main

func majorityElement(nums []int) []int {
	// O1的空间复杂度, 基本就是位运算了吧...

	var limit = len(nums) / 3

	var set = make(map[int]int)

	var ret []int

	const MAX = 5e5

	for _, v := range nums {
		var cnt = set[v] + 1

		if cnt > limit {
			if cnt < MAX {
				ret = append(ret, v)
				set[v] = MAX
			}
		} else {
			set[v] = cnt
		}
	}

	return ret
}

func majorityElementGood(nums []int) (ans []int) {
	element1, element2 := 0, 0
	vote1, vote2 := 0, 0

	for _, num := range nums {
		if vote1 > 0 && num == element1 { // 如果该元素为第一个元素，则计数加1
			vote1++
		} else if vote2 > 0 && num == element2 { // 如果该元素为第二个元素，则计数加1
			vote2++
		} else if vote1 == 0 { // 选择第一个元素
			element1 = num
			vote1++
		} else if vote2 == 0 { // 选择第二个元素
			element2 = num
			vote2++
		} else { // 如果三个元素均不相同，则相互抵消1次
			vote1--
			vote2--
		}
	}

	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if vote1 > 0 && num == element1 {
			cnt1++
		}
		if vote2 > 0 && num == element2 {
			cnt2++
		}
	}

	// 检测元素出现的次数是否满足要求
	if vote1 > 0 && cnt1 > len(nums)/3 {
		ans = append(ans, element1)
	}
	if vote2 > 0 && cnt2 > len(nums)/3 {
		ans = append(ans, element2)
	}
	return
}
