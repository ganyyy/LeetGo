package main

func maxSatisfied(customers []int, grumpy []int, X int) int {

	// 先看一下原始情况下的满意度是多少
	var origin int
	for i := range customers {
		if grumpy[i] == 0 {
			origin += customers[i]
		}
	}

	// 然后把整个区间依次加进去
	for i := 0; i < X; i++ {
		if grumpy[i] == 1 {
			origin += customers[i]
		}
	}

	// 搞区间和
	var res = origin
	for i := X; i < len(customers); i++ {
		if grumpy[i-X] == 1 {
			origin -= customers[i-X]
		}
		if grumpy[i] == 1 {
			origin += customers[i]
		}
		res = max(res, origin)
	}

	return res
}

func maxSatisfied2(customers []int, grumpy []int, minutes int) int {
	// 求和
	var sum int
	// 窗口内通过转换是否生气而增加的收益
	var windowAdd int
	var maxAdd int

	maxAdd = max(maxAdd, windowAdd)
	for i := 0; i < len(customers); i++ {
		if i >= minutes && grumpy[i-minutes] == 1 {
			windowAdd -= customers[i-minutes]
		}
		if grumpy[i] == 1 {
			windowAdd += customers[i]
		} else {
			sum += customers[i]
		}
		maxAdd = max(maxAdd, windowAdd)
	}
	return sum + maxAdd
}
