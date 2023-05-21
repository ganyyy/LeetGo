package main

func addNegabinary(arr1 []int, arr2 []int) (ans []int) {
	i := len(arr1) - 1
	j := len(arr2) - 1
	// 从低位开始加, 有多种可能
	// 0: 低位相加没有产生进位
	// 1: 低位carry=-1, 并且本位都是0, 所以需要本位变成1, 高位进1
	// -1: 低位相加产生了进位, 注意: 这里是-2进制, 所以进位是-1
	carry := 0
	for i >= 0 || j >= 0 || carry != 0 {
		x := carry
		if i >= 0 {
			x += arr1[i]
		}
		if j >= 0 {
			x += arr2[j]
		}
		/*
		    x的可能
		   -1: (0,0)+-1        -> append(1), c=1
		    0: (0,0)+0         -> append(0), c=0
		    1: (0,1)+0/(0,0)+1 -> append(1), c=0
		    2: (1,1)+0/(1,0)+1 -> append(0), c=-1
		    3: (1,1)+1         -> append(1), c=-1
		*/

		if x >= 2 {
			// 2, 3
			ans = append(ans, x-2)
			carry = -1
		} else if x >= 0 {
			// 0, 1
			ans = append(ans, x)
			carry = 0
		} else {
			// -1
			// 为啥呢? 4 = 8+(-4), 当前位保留, 高位进一
			ans = append(ans, 1)
			carry = 1
		}
		i--
		j--
	}
	// 去除前导0
	for len(ans) > 1 && ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}
	// 反转
	for left, n := 0, len(ans); left < n/2; left++ {
		ans[left], ans[n-1-left] = ans[n-1-left], ans[left]
	}
	return ans
}
