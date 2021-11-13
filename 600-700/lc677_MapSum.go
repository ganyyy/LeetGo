package main

type Node struct {
	Next   [26]*Node
	Cnt    int32 // 经过该点为的累加值
	EndCnt int32 // 终点值
}

type MapSum struct {
	Node
}

func Constructor() MapSum {
	return MapSum{}
}

func (ms *MapSum) Add(key string, val int) *Node {
	var cur = &ms.Node
	cur.Cnt += int32(val)
	for i := 0; i < len(key); i++ {
		var next = cur.Next[key[i]-'a']
		if next == nil {
			next = &Node{Cnt: int32(val)}
			cur.Next[key[i]-'a'] = next
		} else {
			next.Cnt += int32(val)
		}
		cur = next
	}
	return cur
}

func (ms *MapSum) Insert(key string, val int) {
	var node = ms.Add(key, val)
	if node.EndCnt != 0 {
		// 这是一个重复键值对, 就去掉差值
		ms.Add(key, -int(node.EndCnt))
	}
	node.EndCnt = int32(val)
}

func (ms *MapSum) Sum(prefix string) int {
	var cur = &ms.Node
	for i := 0; i < len(prefix); i++ {
		cur = cur.Next[prefix[i]-'a']
		if cur == nil {
			break
		}
	}
	if cur != nil {
		return int(cur.Cnt)
	}
	return 0
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
