package main

type Node struct {
	Next  [26]*Node
	IsEnd bool
}

type StreamChecker struct {
	root *Node
	buff []byte
}

func Constructor(words []string) StreamChecker {
	var root = &Node{}

	var insert = func(s string) {
		cur := root
		for i := len(s) - 1; i >= 0; i-- {
			next := cur.Next[s[i]-'a']
			if next == nil {
				next = &Node{}
				cur.Next[s[i]-'a'] = next
			}
			cur = next
		}
		cur.IsEnd = true
	}

	for _, word := range words {
		insert(word)
	}

	return StreamChecker{
		root: root,
	}
}

func (c *StreamChecker) Query(letter byte) bool {

	c.buff = append(c.buff, letter)
	if len(c.buff) > 200 {
		c.buff = c.buff[1:]
	}

	var cur = c.root
	for i := len(c.buff) - 1; i >= 0; i-- {
		idx := int(c.buff[i] - 'a')
		next := cur.Next[idx]
		if next == nil {
			return false
		}
		if next.IsEnd {
			return true
		}
		cur = next
	}
	return cur.IsEnd
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
