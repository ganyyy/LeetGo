package main

import "strconv"

func maximumSwap(num int) int {

	s := []byte(strconv.Itoa(num))
	n := len(s)
	maxIdx, idx1, idx2 := n-1, -1, -1
	// 从后向前, 找到最大位置和小于最大位置的最高位
	for i := n - 1; i >= 0; i-- {
		if s[i] > s[maxIdx] {
			maxIdx = i
		} else if s[i] < s[maxIdx] {
			// maxIdx是最大值对应的位置
			// idx1是小于[maxIdx]的最高位
			idx1, idx2 = i, maxIdx
		}
	}
	if idx1 < 0 {
		// 所有的高位都大于等于低位, 无需交换
		return num
	}
	s[idx1], s[idx2] = s[idx2], s[idx1]
	v, _ := strconv.Atoi(string(s))
	return v
}

func maximumSwap2(num int) int {
	var buffer [10]int
	var numBuffer = buffer[:0]

	var maxNum = num

	for num != 0 {
		numBuffer = append(numBuffer, num%10)
		num /= 10
	}

	// 找到首个最大值(位阶越小越好)
	// 然后正序查找(相较于原数字是逆序的)
	var maxIdx int
	var first, second = 0, -1

	for i := 1; i < len(numBuffer); i++ {
		maxVal := numBuffer[maxIdx]
		curVal := numBuffer[i]
		if maxVal < curVal {
			maxIdx = i
		} else if maxVal > curVal {
			first, second = maxIdx, i
		}
	}

	if second == -1 {
		return maxNum
	}

	numBuffer[first], numBuffer[second] = numBuffer[second], numBuffer[first]

	maxNum = 0
	for i := len(numBuffer) - 1; i >= 0; i-- {
		maxNum = maxNum*10 + numBuffer[i]
	}

	return maxNum
}
