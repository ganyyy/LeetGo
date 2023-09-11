package main

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	g := make([][]int, numCourses)
	indgree := make([]int, numCourses)
	isPre := make([][]bool, numCourses)
	for i, _ := range isPre {
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
