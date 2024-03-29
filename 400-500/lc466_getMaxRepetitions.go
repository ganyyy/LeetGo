package main

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	len1, len2 := len(s1), len(s2)
	index1, index2 := 0, 0 // 注意此处直接使用 Ra Rb 的下标，不取模

	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
		return 0
	}
	// map1: 当出现在Sa末尾时, Sb当前指向的位置 和其对应的 Ra的位置
	// map2: 当出现在Sa末尾时, Sb当前指向的位置 和其对应的 Rb的位置
	map1, map2 := make(map[int]int), make(map[int]int)
	ans := 0 // 注意，此处存储的是 Ra 中 Sb 的个数，而非 Ra 中 Rb 的个数

	for index1/len1 < n1 { // 遍历整个 Ra
		// idx1: 当前迭代的Sa的位置
		// idx2: 当前迭代的Sb的位置
		idx1, idx2 := index1%len1, index2%len2
		if idx1 == len1-1 { // 在 Sa 末尾
			if val, ok := map1[idx2]; ok { // 出现了循环，进行快进
				// 如果Sb中的某个字符两次出现在了Sa的末尾, 就意味着产生了循环节
				cycleLen := index1/len1 - val/len1            // 每个循环占多少个 Sa
				cycleNum := (n1 - 1 - index1/len1) / cycleLen // 还有多少个循环
				cycleS2Num := index2/len2 - map2[idx2]/len2   // 每个循环含有多少个 Sb

				index1 += cycleNum * cycleLen * len1 // 将 Ra 快进到相应的位置
				ans += cycleNum * cycleS2Num         // 把快进部分的答案数量加上
			} else { // 第一次，注意存储的是未取模的
				map1[idx2] = index1
				map2[idx2] = index2
			}

		}

		if s1[idx1] == s2[idx2] {
			if idx2 == len2-1 {
				ans += 1
			}
			// 找到一个匹配的位置, index2+1
			index2 += 1
		}
		// S1会一直增加
		index1 += 1
	}
	return ans / n2
}
