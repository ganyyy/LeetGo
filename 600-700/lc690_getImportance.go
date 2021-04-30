package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	// 搞一个dfs, 然后通过map进行枝减

	var m = make(map[int]*Employee, len(employees))

	for _, e := range employees {
		m[e.Id] = e
	}

	var dfs func(e *Employee)
	var ret int
	dfs = func(e *Employee) {
		if e == nil {
			return
		}
		ret += e.Importance
		delete(m, e.Id)
		for _, i := range e.Subordinates {
			if e, ok := m[i]; ok {
				dfs(e)
			}
		}
	}

	dfs(m[id])

	return ret
}
