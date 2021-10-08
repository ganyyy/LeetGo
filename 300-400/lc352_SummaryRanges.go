package main

type SummaryRanges struct {
	arr [][2]int
}

func ConstructorSummaryRanges() SummaryRanges {
	return SummaryRanges{}
}

func (s *SummaryRanges) AddNum(val int) {
	key := len(s.arr)
	for k, v := range s.arr {
		// 1. 如果新加入的值已经包含在原始的区间中, 直接返回即可
		if v[0] <= val && v[1] >= val {
			return
		}
		if v[0] > val {
			key = k
			break
		}
	}
	if len(s.arr) == 0 {
		// 2. 没有元素的情况下, 当前元素就是第一个
		s.arr = append(s.arr, [2]int{val, val})
		return
	}
	if key == 0 {
		// 小于原始区间的最小值
		left := s.arr[key][0]
		if val+1 == left {
			// 恰好可以组合成一个区间
			s.arr[key][0] = val
		} else {
			// 单独区间
			s.arr = append([][2]int{{val, val}}, s.arr...)
		}
	} else if key == len(s.arr) {
		// 大于原始区间的最大值
		right := s.arr[key-1][1]
		if val-1 == right {
			// 恰好可以组成一个区间
			s.arr[key-1][1] = val
		} else {
			// 单独区间
			s.arr = append(s.arr, [2]int{val, val})
		}
	} else {
		// 获取前置区间右端点和后置区间的左端点
		left, right := s.arr[key-1][1], s.arr[key][0]
		if left+1 == val && right > val+1 {
			// 扩充左边
			s.arr[key-1][1] = val
		} else if left+1 < val && right == val+1 {
			// 扩充右边
			s.arr[key][0] = val
		} else if left+1 == val && val+1 == right {
			// 将左右两边合并
			s.arr[key-1][1] = s.arr[key][1]
			s.arr = append(s.arr[0:key], s.arr[key+1:]...)
		} else {
			// 单独的区间
			s.arr = append(s.arr[:key], append([][2]int{{val, val}}, s.arr[key:]...)...)
		}
	}

}

func (s *SummaryRanges) GetIntervals() [][]int {
	res := make([][]int, len(s.arr))
	for k, v := range s.arr {
		res[k] = []int{v[0], v[1]}
	}
	return res
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
