package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	// 这个不就是按照字典序排序吗....
	// var m = make(map[int]string)

	var get = func(i int) string {
		// var si, ok = m[nums[i]]
		// if !ok {
		//	si = strconv.Itoa(nums[i])
		//	m[nums[i]] = si
		// }
		return strconv.Itoa(nums[i])
	}
	var total int
	sort.Slice(nums, func(i, j int) bool {
		var si, sj = get(i), get(j)
		total += len(si) + len(sj)
		if len(si) == len(sj) {
			return strings.Compare(si, sj) >= 1
		}
		// 判断是否存在相同的前缀
		var hasPrefix bool
		if len(si) > len(sj) {
			hasPrefix = strings.HasPrefix(si, sj)
		} else {
			hasPrefix = strings.HasPrefix(sj, si)
		}
		if hasPrefix {
			return strings.Compare(si+sj, sj+si) >= 1
		}
		return strings.Compare(si, sj) >= 1
	})

	var bs bytes.Buffer
	bs.Grow(total)
	// 特殊处理一下
	if len(nums) > 0 && nums[0] == 0 {
		return "0"
	}
	for _, v := range nums {
		bs.WriteString(strconv.Itoa(v))
	}
	return bs.String()
}

func largestNumber2(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

func largestNumber3(nums []int) string {
	if len(nums) == 0 {
		return "0"
	}
	bitCount := func(n int) int {
		// 找到大于n的首个10的整数次幂
		var bit = 10
		for bit <= n {
			bit *= 10
		}
		return bit
	}

	sort.Slice(nums, func(i, j int) bool {
		n1, n2 := nums[i], nums[j]
		// 计算位数
		bit1, bit2 := bitCount(n1), bitCount(n2)
		return n1*bit2+n2 > n2*bit1+n1
	})

	fmt.Println(nums)

	if nums[0] == 0 {
		return "0"
	}
	var sb strings.Builder
	for _, num := range nums {
		sb.WriteString(strconv.Itoa(num))
	}
	return sb.String()
}

func main() {
	var nums = []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))
}
