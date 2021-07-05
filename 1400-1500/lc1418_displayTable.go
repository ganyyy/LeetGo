package main

import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {

	// 合并同一桌子上的菜品, 和点餐人无关?
	var m = make(map[string]map[string]int)
	var foodM = make(map[string]int)

	var add = func(table, food string) {
		var tm, ok = m[table]
		if !ok {
			tm = make(map[string]int)
			m[table] = tm
		}
		tm[food]++
		foodM[food] = 0
	}

	for _, info := range orders {
		add(info[1], info[2])
	}

	var foodSlice = make([]string, 1, len(foodM)+1)
	foodSlice[0] = "Table"
	for food := range foodM {
		foodSlice = append(foodSlice, food)
	}
	sort.Strings(foodSlice[1:])

	var res = make([][]string, 1, len(m)+1)
	res[0] = foodSlice
	for table, foods := range m {
		var show = make([]string, 1, len(foodSlice))
		show[0] = table

		for _, s := range foodSlice[1:] {
			if n, ok := foods[s]; ok {
				show = append(show, strconv.Itoa(n))
			} else {
				show = append(show, "0")
			}
		}
		res = append(res, show)
	}
	// 坐等超时. 这尼玛还能双百???
	sort.Slice(res[1:], func(i, j int) bool {
		var a, _ = strconv.Atoi(res[i+1][0])
		var b, _ = strconv.Atoi(res[j+1][0])
		return a < b
	})

	return res
}
