package main

import (
	"math"
	"math/bits"
)

func maxStudents(seats [][]byte) int {
	// rowCount: 行的数量
	// colCount: 列的数量
	rowCount, colCount := len(seats), len(seats[0])
	cache := make(map[int]int)

	// 同行是否合法
	isSingleRowCompliant := func(status, row int) bool {
		for col := 0; col < colCount; col++ {
			// 当前位置上尝试坐人
			if (status>>col)&1 == 1 {
				// 位置不合法
				if seats[row][col] == '#' {
					return false
				}
				// 前边有人
				if col > 0 && (status>>(col-1))&1 == 1 {
					return false
				}
			}
		}
		return true
	}

	// 跨行是否合法
	isCrossRowsCompliant := func(status, upperRowStatus int) bool {
		for col := 0; col < colCount; col++ {
			if (status>>col)&1 == 1 {
				// 上一行的左前方
				if col > 0 && (upperRowStatus>>(col-1))&1 == 1 {
					return false
				}
				// 上一行的右前方
				if col < colCount-1 && (upperRowStatus>>(col+1))&1 == 1 {
					return false
				}
			}
		}
		return true
	}

	var dp func(row int, status int) int
	dp = func(row, status int) int {
		// (每一行+状态)构建唯一key
		// 因为status的上限就是2^colCount-1, 所以组成的key就是row*(2^colCount)+status
		statusKey := (row << colCount) + status
		studentNum, find := cache[statusKey]
		if !find {
			// 这一行是否合法
			if !isSingleRowCompliant(status, row) {
				cache[statusKey] = math.MinInt32
				return math.MinInt32
			}
			// 这一行有几个学生
			students := bits.OnesCount(uint(status))
			if row == 0 {
				// 第一行不需要判断前置行
				cache[statusKey] = students
				return students
			}
			var maxStudent int
			// 迭代上一行的状态
			for upperRowStatus := 0; upperRowStatus < 1<<colCount; upperRowStatus++ {
				// 判断一下是不是跨行的
				if isCrossRowsCompliant(status, upperRowStatus) {
					// 从上一行的状态中找到最大的学生数
					maxStudent = max(maxStudent, dp(row-1, upperRowStatus))
				}
			}
			studentNum = students + maxStudent
			cache[statusKey] = studentNum
		}
		return studentNum
	}

	maxStudent := 0
	for status := 0; status < (1 << colCount); status++ {
		// 从最后一行向上推..?
		maxStudent = max(maxStudent, dp(rowCount-1, status))
	}
	return maxStudent
}
