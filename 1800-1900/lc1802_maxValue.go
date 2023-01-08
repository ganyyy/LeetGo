package main

import (
	"fmt"
	"sort"
)

/*
int maxValue(int n, int index, int maxSum) {
    int left = 1, right = maxSum;
    while (left < right) {
        int mid = (left + right + 1) / 2;
        if (valid(mid, n, index, maxSum)) {
            left = mid;
        } else {
            right = mid - 1;
        }
    }
    return left;
}

bool valid(int mid, int n, int index, int maxSum) {
    int left = index;
    int right = n - index - 1;
    return mid + cal(mid, left) + cal(mid, right) <= maxSum;
}

long cal(int big, int length) {
    if (length + 1 < big) {
        int small = big - length;
        return (long) (big - 1 + small) * length / 2;
    } else {
        int ones = length - (big - 1);
        return (long) big * (big - 1) / 2 + ones;
    }
}
*/

func maxValue(n int, index int, maxSum int) int {
	sum := func(v, n int) (ret int) {
		defer func() {
			fmt.Println("ret", v, n, ret)
		}()
		if n+1 < v {
			small := v - n
			return (v - 1 + small) * n / 2
		}
		ones := n - v + 1
		return v*(v-1)/2 + ones
	}
	left := index
	right := n - index - 1
	calc := func(target int) int {
		// nums[index] = target
		// 0.....1...target...1...0
		// 求和
		l := sum(target, left)
		r := sum(target, right)
		fmt.Println("search", target, l, r)
		return l + r + target
	}
	var l = 1
	var r = maxSum
	for l < r {
		// 核心是向上取整?
		var mid = (l + r + 1) / 2
		if calc(mid) <= maxSum {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

func f(big, length int) (ret int) {
	defer func() {
		fmt.Println("ret", big, length, ret)
	}()
	if length == 0 {
		return 0
	}
	if length <= big {
		return (2*big + 1 - length) * length / 2
	}
	return (big+1)*big/2 + length - big
}

func maxValue2(n, index, maxSum int) int {
	// [0...index...n]
	// left  [0, index]
	// right [index+1, n]
	left := index
	right := n - index - 1
	// golang 的 search 是向下取整
	return sort.Search(maxSum, func(mid int) bool {
		l := f(mid, left)
		r := f(mid, right)
		fmt.Println("search", mid, l, r)
		return mid+l+r >= maxSum
	})
}
