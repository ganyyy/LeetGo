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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
