package main

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	var ln = len(nums)
	var l, r = 0, ln

	for l < r {
		var mid = l + ((r - l) >> 1)

		// 判断左右剩余的数字

		// 先看是不是就这一个数
		var find, isPre bool
		if mid >= 1 && nums[mid-1] == nums[mid] {
			mid--
			isPre = true
			find = true
		} else if mid < ln-1 && nums[mid+1] == nums[mid] {
			find = true
		}
		// fmt.Println(l, mid, r)

		// 没找到, 那就说明只存在一次的数
		if !find {
			return nums[mid]
		}

		// 判断向左还是向右

		// 判断从数开始的位置开始到结尾, 如果是偶数个数字(差值为奇数), 就说明在左边, 否则就是在右边
		if (ln-mid)&1 == 0 {
			// 右边剩余偶数, 那么一定在左边
			r = mid
		} else {
			if isPre {
				mid++
			}
			l = mid + 1
		}

	}

	return nums[r]
}

/*
func singleNonDuplicate(nums []int) int {
    var ret int

    for _, v := range nums {
        ret ^= v
    }

    return ret
}
*/
