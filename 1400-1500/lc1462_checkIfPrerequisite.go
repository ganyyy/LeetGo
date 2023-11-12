package main

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	g := make([][]int, numCourses)
	indgree := make([]int, numCourses)
	isPre := make([][]bool, numCourses)
	for i := range isPre {
		isPre[i] = make([]bool, numCourses)
		g[i] = []int{}
	}
	for _, p := range prerequisites {
		indgree[p[1]]++
		g[p[0]] = append(g[p[0]], p[1])
	}

	var q, next []int
	for i := 0; i < numCourses; i++ {
		if indgree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		for _, cur := range q {
			for _, ne := range g[cur] {
				// 这个建立映射图的逻辑很关键
				isPre[cur][ne] = true
				for i := 0; i < numCourses; i++ {
					isPre[i][ne] = isPre[i][ne] || isPre[i][cur]
				}
				indgree[ne]--
				if indgree[ne] == 0 {
					next = append(next, ne)
				}
			}
		}
		q, next = next, q[:0]
	}
	var res = make([]bool, 0, len(queries))
	for _, query := range queries {
		res = append(res, isPre[query[0]][query[1]])
	}
	return res
}

func checkIfPrerequisite2(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// TODO 查询缓存优化
	var preCourses = make([][]int, numCourses)
	for _, prerequisite := range prerequisites {
		preCourses[prerequisite[1]] = append(preCourses[prerequisite[1]], prerequisite[0])
	}

	var visit = make([]bool, numCourses)

	var matchRoot func(course, pre int) bool
	matchRoot = func(course, pre int) bool {
		if course == pre {
			return true
		}
		for _, preCourse := range preCourses[course] {
			if preCourse == pre {
				return true
			}
			if visit[preCourse] {
				continue
			}
			visit[preCourse] = true
			if matchRoot(preCourse, pre) {
				return true
			}
		}
		return false
	}

	var ret = make([]bool, 0, len(queries))

	for _, query := range queries {
		need, course := query[0], query[1]
		visit[course] = true
		ret = append(ret, matchRoot(course, need))
		for i := range visit {
			visit[i] = false
		}
	}

	return ret
}

func checkIfPrerequisite3(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	// graph[i] 表示 i的后置课程
	graph := make([][]int, numCourses)
	// inDegree[i] 表示 i的入度, 即i还有多少前置课程未完成
	inDegree := make([]int, numCourses)
	// isPreCourse[i][j] 表示 i是j的前置课程
	isPreCourse := make([][]bool, numCourses)
	for course := range isPreCourse {
		isPreCourse[course] = make([]bool, numCourses)
		graph[course] = []int{}
	}
	for _, prerequisite := range prerequisites {
		pre, next := prerequisite[0], prerequisite[1]
		// next入度+1
		inDegree[next]++
		// pre的后置课程有next
		graph[pre] = append(graph[pre], next)
	}

	var curQueue, nextQueue []int
	for course := 0; course < numCourses; course++ {
		if inDegree[course] == 0 {
			// 入度为0, 说明没有前置课程, 可以直接学习
			curQueue = append(curQueue, course)
		}
	}

	for len(curQueue) > 0 {
		for _, cur := range curQueue {
			for _, next := range graph[cur] {
				isPreCourse[cur][next] = true
				for i := 0; i < numCourses; i++ {
					// i是否是ne的前置课程呢? 主要由两个条件决定:
					// i是ne的直接前置
					// i是ne的间接前置(i是cur的前置)
					isPreCourse[i][next] = isPreCourse[i][next] || isPreCourse[i][cur]
				}
				inDegree[next]--
				if inDegree[next] == 0 {
					// next的入度为0, 说明next的前置课程都已经完成, 可以继续学习
					nextQueue = append(nextQueue, next)
				}
			}
		}
		curQueue, nextQueue = nextQueue, curQueue[:0]
	}
	var res = make([]bool, 0, len(queries))
	for _, query := range queries {
		res = append(res, isPreCourse[query[0]][query[1]])
	}
	return res
}
