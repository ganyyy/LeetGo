package main

import "fmt"

func lengthLongestPath(input string) int {
	var depth [30]int // 每一级的最大长度
	var res int
	input += "\n" // 强制遍历到最后一位
	// 标记是不是文件
	// 当前等级
	// 当前文件(夹)长度
	isFile, level, cur := false, 0, 0
	for i := 0; i < len(input); i++ {
		c := input[i]
		// 如果是换行, 就将level归零
		if c == '\n' {
			if level > 0 {
				depth[level] = depth[level-1] + cur + 1
			} else {
				depth[level] = cur
			}
			if isFile {
				res = max(res, depth[level])
				isFile = false
			}
			cur = 0
			level = 0
			continue
		}
		if c == '\t' {
			level++
			continue
		}
		if c == '.' {
			isFile = true
		}
		cur++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthLongestPath2(input string) int {
	var stack []int

	stack = append(stack, 0)
	var i int
	var maxLen int
	var max = func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	var curDepth int
	for i < len(input) {
		if input[i] == '\n' {
			curDepth = 0
			for i = i + 1; i < len(input) && input[i] == '\t'; i++ {
				curDepth++
			}
			if curDepth >= len(stack) { // 只能一级一级向下增长
				stack = append(stack, 0)
			}
			if curDepth == 0 {
				stack[curDepth] = 0 // 更换新的根目录, 需要重置
				continue
			}
			stack[curDepth] = stack[curDepth-1] + 1 // 保留父级目录的长度
		} else {
			var isFile bool
			var start = i
			for ; i < len(input) && input[i] != '\n'; i++ {
				if input[i] == '.' {
					isFile = true
				}
			}
			var name = input[start:i]
			if isFile {
				maxLen = max(maxLen, len(name)+stack[curDepth])
			} else {
				stack[curDepth] += len(name)
			}
		}
	}

	return maxLen
}

func main() {
	fmt.Println(lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"))
}
