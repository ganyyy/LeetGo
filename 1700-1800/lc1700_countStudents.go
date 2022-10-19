package main

func countStudents(students, sandwiches []int) int {
	// 队列中0/1的计数
	s1 := 0
	for _, v := range students {
		s1 += v
	}
	s0 := len(students) - s1
	for _, x := range sandwiches {
		// 和学生相对位置无关, 有就可以用
		// 只看队列的首个值, 是否有需求
		// 没有学生可以消费, 那么剩下的学生也都无法消费, 直接跳出即可
		// **因为只能从开头消费**

		if x == 0 && s0 > 0 {
			s0--
		} else if x == 1 && s1 > 0 {
			s1--
		} else {
			break
		}
	}
	// 返回剩余的个数
	return s0 + s1
}
