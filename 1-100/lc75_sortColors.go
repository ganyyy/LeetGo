package main

import "fmt"

func sortColors(nums []int) {
	n1, n2, n3 := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case 0:
			n1++
		case 1:
			n2++
		case 2:
			n3++
		}
	}
	index := 0
	addVal := func(count, val int) {
		for i := 0; i < count; i++ {
			nums[index] = val
			index++
		}
	}
	addVal(n1, 0)
	addVal(n2, 1)
	addVal(n3, 2)
}

func sortColors2(nums []int) {
	// 分别代表每个数字的起始索引
	zero, one, two := -1, 0, len(nums)
	for one < two {
		if nums[one] == 0 {
			// 如果 one = 0, 那就把他换给zero的下一个位置
			zero++
			nums[zero], nums[one] = nums[one], nums[zero]
			one++
		} else if nums[one] == 2 {
			// 如果 one == 2, 那就把他换给two的前一个位置
			two--
			nums[two], nums[one] = nums[one], nums[two]
		} else {
			// 如果 one == 1, 直接加呗
			one++
		}
	}
}

func sortColors3(nums []int) {
	// 分别代表每个数字的起始索引
	_0, _1, _2 := -1, 0, len(nums)
	for _1 < _2 {
		switch nums[_1] {
		case 0:
			// 如果是0的话, 移动到浅表
			_0++
			nums[_0], nums[_1] = nums[_1], nums[_0]
			_1++
		case 1:
			// 如果是1的话, 跳过看下一个
			_1++
		case 2:
			// 如果是2的话, 移动到后边
			_2--
			nums[_2], nums[_1] = nums[_1], nums[_2]
		}
	}
}

func main() {
	v := []int{
		1, 2, 0,
	}
	sortColors2(v)
	fmt.Println(v)
}
